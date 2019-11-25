package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/balancer"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/delivery"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	repositoryMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase/middleware"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	customMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pkg/middlewares"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	consulAddr = flag.String("addr", "127.0.0.1:8500", "consul addr (8500 in original consul)")
)

var (
	consul       *consulapi.Client
	nameResolver *balancer.TestNameResolver
)


func AuthServiceCreate() (auth services.AuthorizationServiceClient) {
	var err error
	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	consul, err = consulapi.NewClient(config)

	health, _, err := consul.Health().Service("authorization-service", "", false, nil)
	if err != nil {
		log.Fatalf("cant get alive services")
	}

	servers := []string{}
	for _, item := range health {
		addr := item.Service.Address +
			":" + strconv.Itoa(item.Service.Port)
		servers = append(servers, addr)
	}

	nameResolver = &balancer.TestNameResolver{
		Addr: servers[0],
	}

	grcpConn, err := grpc.Dial(
		servers[0],
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBalancer(grpc.RoundRobin(nameResolver)),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	if len(servers) > 1 {
		var updates []*naming.Update
		for i := 1; i < len(servers); i++ {
			updates = append(updates, &naming.Update{
				Op:   naming.Add,
				Addr: servers[i],
			})
		}
		nameResolver.W.Inject(updates)
	}

	sessManager := services.NewAuthorizationServiceClient(grcpConn)

	// тут мы будем периодически опрашивать консул на предмет изменений
	go runOnlineServiceDiscovery(servers)

	testUserReg := services.UserReg{
		Email:                "test@mail.ru",
		Password:             "12314ghMEFnk123",
		Username:             "Test",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	serviceCtx  := context.Background()

	cookie, err := sessManager.RegUser(serviceCtx, &testUserReg)
	println(cookie, err)


	return sessManager
}

func runOnlineServiceDiscovery(servers []string) {
	currAddrs := make(map[string]struct{}, len(servers))
	for _, addr := range servers {
		currAddrs[addr] = struct{}{}
	}
	ticker := time.Tick(5 * time.Second)
	for _ = range ticker {
		health, _, err := consul.Health().Service("authorization-service", "", false, nil)
		if err != nil {
			log.Fatalf("cant get alive services")
		}

		newAddrs := make(map[string]struct{}, len(health))
		for _, item := range health {
			addr := item.Service.Address +
				":" + strconv.Itoa(item.Service.Port)
			newAddrs[addr] = struct{}{}
		}

		var updates []*naming.Update
		// проверяем что удалилось
		for addr := range currAddrs {
			if _, exist := newAddrs[addr]; !exist {
				updates = append(updates, &naming.Update{
					Op:   naming.Delete,
					Addr: addr,
				})
				delete(currAddrs, addr)
				fmt.Println("remove", addr)
			}
		}
		// проверяем что добавилось
		for addr := range newAddrs {
			if _, exist := currAddrs[addr]; !exist {
				updates = append(updates, &naming.Update{
					Op:   naming.Add,
					Addr: addr,
				})
				currAddrs[addr] = struct{}{}
				fmt.Println("add", addr)
			}
		}
		if len(updates) > 0 {
			nameResolver.W.Inject(updates)
		}
	}
}

func main() {

	auth := AuthServiceCreate()

	e := echo.New()
	middlewares := customMiddleware.MiddlewareStruct{}
	mRep := repositoryMiddleware.MRepositoryStruct{}
	err := mRep.DataBaseInit()
	if err != nil {
		e.Logger.Error("can't connect to database " + err.Error())
		return
	}
	mUseCase := useCaseMiddleware.MUseCaseStruct{}
	mUseCase.NewUseCase(&mRep)
	middlewares.NewMiddleware(e, &mUseCase, auth)
	//e.Use(customMiddleware.CORSMiddleware)
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	//e.Use(customMiddleware.PanicMiddleware)
	//e.Use(customMiddleware.AccessLogMiddleware)
	//e.Use(customMiddleware.AuthenticationMiddleware)
	//e.HTTPErrorHandler = customMiddleware.CustomHTTPErrorHandler

	e.Static("/static", "static")

	handlers := delivery.HandlersStruct{}
	var mutex sync.Mutex
	rep := repository.ReposStruct{}
	err = rep.DataBaseInit()
	if err != nil {
		e.Logger.Error("can't connect to database " + err.Error())
		return
	}
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	useCase := usecase.UseStruct{}
	useCase.NewUseCase(&mutex, &rep, &san, hub)
	err = handlers.NewHandlers(e, &useCase, auth)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	if err := e.Start(consts.HostAddress); err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")
}

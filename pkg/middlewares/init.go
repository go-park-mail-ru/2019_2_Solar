package middlewares

import (
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	useCaseMiddleware "github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase/middleware"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/consts"
	echov4 "github.com/labstack/echo/v4"
	//"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
)

//var (
//	consulAddr = flag.String("addr", "127.0.0.1:8500", "consul addr (8500 in original consul)")
//)
//
//var (
//	consul       *consulapi.Client
//	nameResolver *balancer.TestNameResolver
//)


func (MS *MiddlewareStruct) NewMiddleware(e *echov4.Echo, mRep useCaseMiddleware.MUseCaseInterface, auth services.AuthorizationServiceClient) {
	MS.MUsecase = mRep

	MS.MAuth = auth

	e.Use(MS.CORSMiddleware)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	e.Use(MS.PanicMiddleware)
	e.Use(MS.AuthenticationMiddleware)
	e.HTTPErrorHandler = MS.CustomHTTPErrorHandler

	//var err error
	//config := consulapi.DefaultConfig()
	//config.Address = *consulAddr
	//consul, err = consulapi.NewClient(config)
	//
	//health, _, err := consul.Health().Service("authorization-service", "", false, nil)
	//if err != nil {
	//	log.Fatalf("cant get alive services")
	//}
	//
	//servers := []string{}
	//for _, item := range health {
	//	addr := item.Service.Address +
	//		":" + strconv.Itoa(item.Service.Port)
	//	servers = append(servers, addr)
	//}
	//
	//nameResolver = &balancer.TestNameResolver{
	//	Addr: servers[0],
	//}
	//
	//grcpConn, err := grpc.Dial(
	//	servers[0],
	//	grpc.WithInsecure(),
	//	grpc.WithBlock(),
	//	grpc.WithBalancer(grpc.RoundRobin(nameResolver)),
	//)
	//if err != nil {
	//	log.Fatalf("cant connect to grpc")
	//}
	//defer grcpConn.Close()
	//
	//if len(servers) > 1 {
	//	var updates []*naming.Update
	//	for i := 1; i < len(servers); i++ {
	//		updates = append(updates, &naming.Update{
	//			Op:   naming.Add,
	//			Addr: servers[i],
	//		})
	//	}
	//	nameResolver.W.Inject(updates)
	//}
	//
	//MS.MAuth = services.NewAuthorizationServiceClient(grcpConn)
	//
	//// тут мы будем периодически опрашивать консул на предмет изменений
	//go functions.RunOnlineServiceDiscovery(servers)
	//
	//testUserReg := services.UserReg{
	//	Email:                "test@mail.ru",
	//	Password:             "12314ghMEFnk123",
	//	Username:             "Test",
	//	XXX_NoUnkeyedLiteral: struct{}{},
	//	XXX_unrecognized:     nil,
	//	XXX_sizecache:        0,
	//}
	//
	//serviceCtx  := context.Background()
	//
	//cookie, err := MS.MAuth.RegUser(serviceCtx, &testUserReg)
	//println(cookie, err)
}

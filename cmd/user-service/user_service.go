package main

import (
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/user-service/pinterest/delivery"
	user_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/user-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"sync"
)

var (
	grpcPort   = flag.Int("grpc", 8087, "listen addr")
	consulAddr = flag.String("consul", "127.0.0.1:8500", "consul addr (8500 in original consul)")
)


func main() {
	flag.Parse()

	port := strconv.Itoa(*grpcPort)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()

	var mutex sync.Mutex
	rep := repository.ReposStruct{}
	err = rep.DataBaseInit()
	if err != nil {
		fmt.Println("can't connect to database " + err.Error())
		return
	}
	san := sanitizer.SanitStruct{}
	san.NewSanitizer()
	hub := webSocket.HubStruct{}
	hub.NewHub()

	useCase := usecase.UseStruct{}
	useCase.NewUseCase(&mutex, &rep, &san, hub)
	user_service.RegisterUserServiceServer(server,
		delivery.NewUserService(&useCase, port))

	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	consul, err := consulapi.NewClient(config)

	serviceID := "SAPI_127.0.0.1:" + port

	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "user-service",
		Port:    *grpcPort,
		Address: "127.0.0.1",
	})
	if err != nil {
		fmt.Println("cant add service to consul", err)
		return
	}
	fmt.Println("registered in consul", serviceID)

	defer func() {
		err := consul.Agent().ServiceDeregister(serviceID)
		if err != nil {
			fmt.Println("cant add service to consul", err)
			return
		}
		fmt.Println("deregistered in consul", serviceID)
	}()

	fmt.Println("starting server at " + port)
	go server.Serve(lis)

	fmt.Println("Press any key to exit")
	fmt.Scanln()
}
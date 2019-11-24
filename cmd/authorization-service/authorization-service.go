package main

import (
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/authorization-service/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/repository"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"sync"
)

var (
	grpcPort   = flag.Int("grpc", 8081, "listen addr")
	consulAddr = flag.String("consul", "127.0.0.1:8500", "consul addr (8500 in original consul)")
)

/*
	go run *.go --grpc="8081" --consul="192.168.99.100:32769"
	go run *.go --grpc="8082" --consul="192.168.99.100:32769"
*/

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


	services.RegisterAuthorizationServiceServer(server,
		usecase.NewAuthorizationService(&mutex, &rep, &san, port))

	config := consulapi.DefaultConfig()
	config.Address = *consulAddr
	consul, err := consulapi.NewClient(config)

	serviceID := "SAPI_127.0.0.1:" + port

	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "authorization-service",
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
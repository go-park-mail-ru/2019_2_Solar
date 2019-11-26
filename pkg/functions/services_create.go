package functions


import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/balancer"
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"log"
	"strconv"
	"time"
)

func ServiceCreate(serviceName string) (Client pinboard_service.PinBoardServiceClient) {
	var err error
	config := consulapi.DefaultConfig()
	config.Address = consulAddr
	consul, err = consulapi.NewClient(config)

	health, _, err := consul.Health().Service(serviceName, "", false, nil)
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

	grpcConn, err := grpc.Dial(
		servers[0],
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBalancer(grpc.RoundRobin(nameResolver)),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}

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

	client := pinboard_service.NewPinBoardServiceClient(grpcConn)

	// тут мы будем периодически опрашивать консул на предмет изменений
	go MyOnlineServiceDiscovery(servers, serviceName)

	return client
}

func MyOnlineServiceDiscovery(servers []string, serviceName string) {
	currAddrs := make(map[string]struct{}, len(servers))
	for _, addr := range servers {
		currAddrs[addr] = struct{}{}
	}
	ticker := time.Tick(5 * time.Second)
	for _ = range ticker {
		health, _, err := consul.Health().Service(serviceName, "", false, nil)
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


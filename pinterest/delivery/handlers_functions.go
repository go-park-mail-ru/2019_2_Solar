package delivery

import (
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/labstack/echo"
)

//var (
//	consulAddr = flag.String("addr", "127.0.0.1:8500", "consul addr (8500 in original consul)")
//)
//
//var (
//	consul       *consulapi.Client
//	nameResolver *balancer.TestNameResolver
//)

func (h *HandlersStruct) NewHandlers(e *echo.Echo, useCase usecase.UseInterface, auth services.AuthorizationServiceClient, pinBoardService pinboard_service.PinBoardServiceClient) error {
	h.PUsecase = useCase

	h.AuthSessManager = auth
	h.PinBoardService = pinBoardService


	e.GET("/", h.HandleEmpty)

	e.GET("/users", h.HandleListUsers)
	e.GET("/users/:username", h.HandleGetUserByUsername)

	e.POST("/subscribe/:username", h.HandleCreateSubscribe)
	e.DELETE("/subscribe/:username", h.HandleDeleteSubscribe)

	// ==============================================================

	e.POST("/registration", h.HandleRegUser)
	e.POST("/login", h.HandleLoginUser)
	e.POST("/logout", h.HandleLogoutUser)

	e.POST("/service/registration", h.ServiceRegUser)
	e.POST("/service/login", h.ServiceLoginUser)
	e.POST( "/service/logout", h.ServiceLogoutUser)

	// ==============================================================

	e.GET("/profile/data", h.HandleGetProfileUserData)

	e.POST("/profile/data", h.HandleEditProfileUserData)
	e.POST("/profile/picture", h.HandleEditProfileUserPicture)

	// ==============================================================

	e.POST("/board", h.HandleCreateBoard)
	e.GET("/board/:id", h.HandleGetBoard)

	e.POST("/pin", h.HandleCreatePin)
	e.GET("/pin/:id", h.HandleGetPin)


	e.POST("/service/board", h.ServiceCreateBoard)
	e.GET( "/service/board/:id", h.ServiceGetBoard)

	e.POST("/service/pin", h.ServiceCreatePin)
	e.GET("/service/pin/:id", h.ServiceGetPin)

	// ==============================================================

	e.GET("/board/list/my", h.HandleGetMyBoards)

	e.POST("/pin/:id/comment", h.HandleCreateComment)
	e.GET("/pin/list/new", h.HandleGetNewPins)
	e.GET("/pin/list/my", h.HandleGetMyPins)
	e.GET("/pin/list/subscribe", h.HandleGetSubscribePins)

	e.POST("/notice/:receiver_id", h.HandleCreateNotice)
	e.GET( "/notice", h.HandleGetNotices)

	e.GET("/chat", h.HandleUpgradeWebSocket)

	e.GET( "/find/pins/by/tag/:tag", h.HandlerFindPinByTag)
	e.GET( "/find/users/by/username/:username", h.HandlerFindUserByUsername)

	e.GET ("/categories", h.HandleGetCategories)


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
	//h.AuthSessManager= services.NewAuthorizationServiceClient(grcpConn)
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
	//cookie, err := h.AuthSessManager.RegUser(serviceCtx, &testUserReg)
	//println(cookie, err)

	return nil
}

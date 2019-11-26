package delivery

import (
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
)

type HandlersStruct struct {
	PUsecase usecase.UseInterface
	AuthSessManager services.AuthorizationServiceClient
	PinBoardService pinboard_service.PinBoardServiceClient
}
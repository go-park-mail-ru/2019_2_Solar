package delivery

import (
	pinboard_service "github.com/go-park-mail-ru/2019_2_Solar/cmd/pinboard-service/service_model"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
)

type HandlersStruct struct {
	PUsecase usecase.UseInterface
	AuthSessManager functions.Auth
	PinBoardService pinboard_service.PinBoardServiceClient
}
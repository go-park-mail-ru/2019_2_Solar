package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/cmd/services"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
)

type HandlersStruct struct {
	PUsecase usecase.UseInterface
	AuthSessManager services.AuthorizationServiceClient
}

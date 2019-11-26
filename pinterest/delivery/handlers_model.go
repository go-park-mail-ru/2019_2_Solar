package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/functions"
)

type HandlersStruct struct {
	PUsecase usecase.UseInterface
	AuthSessManager functions.Auth
}

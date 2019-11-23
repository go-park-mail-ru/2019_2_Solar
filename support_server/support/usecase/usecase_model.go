package usecase

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/sanitizer"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Solar/support_server/support/repository"
	webSocket "github.com/go-park-mail-ru/2019_2_Solar/support_server/support/web_socket"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"sync"
)


type UseStruct struct {
	PRepository repository.ReposInterface
	Hub         webSocket.HubStruct
	Mu          *sync.Mutex
}

type UseInterface interface {

}

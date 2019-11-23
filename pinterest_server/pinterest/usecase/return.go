package usecase

import webSocket "github.com/go-park-mail-ru/2019_2_Solar/pinterest_server/pinterest/web_socket"

func (USC *UseStruct) ReturnHub() *webSocket.HubStruct {
	return &USC.Hub
}

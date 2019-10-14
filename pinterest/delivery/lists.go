package delivery

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	h.Mu.Lock()
	data := h.PUsecase.SetJsonData(h.Users, "OK")
	h.Mu.Unlock()

	err := encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while marshalling JSON", err)
		return
	}
}

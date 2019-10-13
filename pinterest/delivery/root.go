package delivery

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handlers) HandleEmpty(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	data := h.PUsecase.SetJsonData(nil, "Empty handler has been done")
	encoder.Encode(data)
	log.Printf("Empty handler has been done")
}

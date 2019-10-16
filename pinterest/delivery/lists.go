package delivery

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func (h *HandlersStruct) HandleListUsers(ctx echo.Context) error {
	w := ctx.Response()

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	users := h.PUsecase.GetAllUsers()
	data := h.PUsecase.SetJsonData(users, "OK")

	err := encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.PUsecase.SetResponseError(encoder, "error while marshalling JSON", err)
		return nil
	}
	return nil
}

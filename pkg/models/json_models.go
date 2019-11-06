package models

type JSONResponse struct {
	Body interface{} `json:"Body"`
}

type JSONInfoResponse struct {
	Body struct {
		Info string `json:"Info"`
	} `json:"Body"`
}

package models

type JSONPin struct {
	Pin interface{} `json:"pin"`
}

type JSONResponse struct {
	Body interface{} `json:"Body"`
}

type JSONInfoResponse struct {
	Body struct {
		Info string `json:"info"`
	} `json:"Body"`
}

type BodyInfo struct {
	Info string `json:"Info"`
}

type ValeraJSONResponse struct {
	CSRF string      `json:"csrf_token"`
	Body interface{} `json:"body"`
}

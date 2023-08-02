package controller

import (
	"encoding/json"
	"net/http"

	"randapi.com/service"
)

type controller struct{}

func NewController() Controller {
	return &controller{}
}

var (
	serviceInstance service.Service = service.NewService()
)

func (*controller) GetRoot(resp http.ResponseWriter, _ *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := serviceInstance.GetRow()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error getting language row"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}

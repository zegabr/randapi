package controller

import (
	"net/http"

	"randapi.com/service"
)

type controller struct{}

var (
	serviceInstance service.Service
)

func NewController(service service.Service) Controller {
	serviceInstance = service
	return &controller{}
}

func (*controller) GetRoot(resp http.ResponseWriter, _ *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := serviceInstance.GetRow()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error getting language row"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(result.Phrase))
}

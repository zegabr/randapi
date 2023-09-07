package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"randapi.com/entity"
	router "randapi.com/http"
	repository "randapi.com/repo"
	"randapi.com/service"
)

var (
	repo       repository.Repository = repository.NewMysqlRepository()
	serv       service.Service       = service.NewService(repo)
	ctrl       Controller            = NewController(serv)
	httpRouter router.Router         = router.NewMuxRouter()
)

func TestGetRoot(t *testing.T) {
	// create the http get request
    var jsonreq io.Reader = nil
	req, _ := http.NewRequest("GET", "/", jsonreq)
	// assign http handler function
    handler := http.HandlerFunc(ctrl.GetRoot)

	// record http response
	response := httptest.NewRecorder()

	// dispatch http request
	handler.ServeHTTP(response, req)

	//add assertions
	status := response.Code
	if status != http.StatusOK {
		t.Error("not returned ok")
	}

	// decode http response
	var row entity.Row
	json.NewDecoder(io.Reader(response.Body)).Decode(&row)
    assert.NotNil(t,row)
}

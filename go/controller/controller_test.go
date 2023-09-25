package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"randapi.com/entity"
	router "randapi.com/http"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) GetRow() (*entity.Row, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Row), args.Error(1)
}

var (
	mockService *MockService   = new(MockService)
	ctrl        Controller    = NewController(mockService)
	httpRouter  router.Router = router.NewMuxRouter()
)

func TestGetRoot(t *testing.T) {

	row := entity.Row{Language: "go", Phrase: "hello from go"}
	mockService.On("GetRow").Return(&row, nil)

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
	json.NewDecoder(io.Reader(response.Body)).Decode(&row)
	assert.NotNil(t, row)
}

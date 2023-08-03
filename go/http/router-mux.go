package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) SERVE(port string) {
	fmt.Println("server listening on port", port)
	http.ListenAndServe(port, muxDispatcher)
}

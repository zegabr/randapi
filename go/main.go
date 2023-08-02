package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"randapi.com/controller"
)

var (
	controllerInstance controller.Controller = controller.NewController()
)

func main() {
	var router = mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", controllerInstance.GetRoot).Methods("GET")

	fmt.Println("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

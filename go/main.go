package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var router = mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", getRoot).Methods("GET")

	fmt.Println("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

package main

import "net/http"

func getRoot(resp http.ResponseWriter, _ *http.Request) {
	result := []byte("hello from Go")
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

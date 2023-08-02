package controller

import "net/http"

type Controller interface {
	GetRoot(resp http.ResponseWriter, req *http.Request)
}

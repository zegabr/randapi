package main

import (
	"randapi.com/controller"
	router "randapi.com/http"
)

var (
	httpRouter         router.Router         = router.NewMuxRouter()
	controllerInstance controller.Controller = controller.NewController()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", controllerInstance.GetRoot)
	httpRouter.SERVE(port)
}

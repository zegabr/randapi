package main

import (
	"os"

	"randapi.com/controller"
	router "randapi.com/http"
	repository "randapi.com/repo"
	"randapi.com/service"
)

var (
	repositoryInstance repository.Repository = repository.NewMysqlRepository()
	serviceInstance    service.Service       = service.NewService(repositoryInstance)
	controllerInstance controller.Controller = controller.NewController(serviceInstance)
	httpRouter         router.Router         = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/", controllerInstance.GetRoot)
	httpRouter.SERVE(os.Getenv("PORT"))
}

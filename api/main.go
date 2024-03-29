package main

import (
	"cow_back/api/dependencies"
	"cow_back/api/server"

	"github.com/gin-gonic/gin"
)

func main() {
	dep := dependencies.BuildDependencies()
	if err := dep.Invoke(func(router *server.Router, server *gin.Engine) {
		router.Resource(server)
		server.Run()
	}); err != nil {
		panic(err)
	}
}

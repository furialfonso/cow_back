package main

import (
	"shared-wallet-service/api/dependencies"
	"shared-wallet-service/api/server"

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

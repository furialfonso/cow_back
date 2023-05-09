package server

import (
	"docker-go-project/api/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pingHandler handlers.IPingHandler
	airHandler  handlers.IAirHandler
}

func NewRouter(pingHandler handlers.IPingHandler,
	airHandler handlers.IAirHandler) *Router {
	return &Router{
		pingHandler,
		airHandler,
	}
}

func (r Router) Resource(gin *gin.Engine) {
	gin.GET("/ping", r.pingHandler.Ping)
	gin.GET("/get-air", r.airHandler.GetAir)
}

package server

import (
	"docker-go-project/api/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pingHandler  handlers.IPingHandler
	groupHandler handlers.IGroupHandler
}

func NewRouter(pingHandler handlers.IPingHandler,
	groupHandler handlers.IGroupHandler) *Router {
	return &Router{
		pingHandler,
		groupHandler,
	}
}

func (r Router) Resource(gin *gin.Engine) {
	gin.GET("/ping", r.pingHandler.Ping)
	gin.POST("/create-group", r.groupHandler.Create)
	gin.GET("/groups", r.groupHandler.GetGroups)
	gin.GET("/group/:code", r.groupHandler.GetGroupByCode)
}

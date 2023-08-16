package server

import (
	"docker-go-project/api/handlers"
	groupHandler "docker-go-project/api/handlers/group"
	userHandler "docker-go-project/api/handlers/user"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pingHandler  handlers.IPingHandler
	groupHandler groupHandler.IGroupHandler
	userHandler  userHandler.IUserHandler
}

func NewRouter(pingHandler handlers.IPingHandler,
	groupHandler groupHandler.IGroupHandler,
	userHandler userHandler.IUserHandler) *Router {
	return &Router{
		pingHandler,
		groupHandler,
		userHandler,
	}
}

func (r Router) Resource(gin *gin.Engine) {
	gin.GET("/ping", r.pingHandler.Ping)
	group := gin.Group("/groups")
	{
		group.GET("", r.groupHandler.GetAll)
		group.GET("/:code", r.groupHandler.GetByCode)
		group.POST("", r.groupHandler.Create)
		group.DELETE("/:code", r.groupHandler.Delete)
	}

	user := gin.Group("/users")
	{
		user.GET("", r.userHandler.GetAll)
		user.GET("/:code", r.userHandler.GetByNickName)
		user.POST("", r.userHandler.Create)
		user.DELETE("/:code", r.userHandler.Delete)
	}
}

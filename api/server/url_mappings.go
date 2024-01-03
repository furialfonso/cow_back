package server

import (
	"docker-go-project/api/handlers"
	groupHandler "docker-go-project/api/handlers/group"
	teamHandler "docker-go-project/api/handlers/team"
	"docker-go-project/api/jobs"
	"sync"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pingHandler  handlers.IPingHandler
	groupHandler groupHandler.IGroupHandler
	teamHandler  teamHandler.ITeamHandler
	job          jobs.IJob
}

func NewRouter(pingHandler handlers.IPingHandler,
	groupHandler groupHandler.IGroupHandler,
	teamHandler teamHandler.ITeamHandler,
	job jobs.IJob) *Router {
	return &Router{
		pingHandler,
		groupHandler,
		teamHandler,
		job,
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

	team := gin.Group("/team")
	{
		team.GET("/:code", r.teamHandler.GetUsersByGroup)
		team.POST("/:code", r.teamHandler.ComposeTeam)
		team.DELETE("/:code", r.teamHandler.DecomposeTeam)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		r.job.UserCache()
	}()
	wg.Wait()
}

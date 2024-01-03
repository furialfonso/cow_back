package server

import (
	"cow_back/api/handlers"
	groupHandler "cow_back/api/handlers/group"
	teamHandler "cow_back/api/handlers/team"
	"cow_back/api/jobs"
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

	team := gin.Group("/teams")
	{
		team.GET("/:code", r.teamHandler.GetTeamByGroup)
		team.GET("/user/:userID", r.teamHandler.GetTeamsByUser)
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

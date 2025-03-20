package server

import (
	"sync"

	"shared-wallet-service/infrastructure/jobs"
	"shared-wallet-service/interfaces/handlers"
	budgetHandler "shared-wallet-service/interfaces/handlers/budget"
	teamHandler "shared-wallet-service/interfaces/handlers/team"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pingHandler   handlers.IPingHandler
	budgetHandler budgetHandler.IBudgetHandler
	teamHandler   teamHandler.ITeamHandler
	job           jobs.IJob
}

func NewRouter(pingHandler handlers.IPingHandler,
	budgetHandler budgetHandler.IBudgetHandler,
	teamHandler teamHandler.ITeamHandler,
	job jobs.IJob,
) *Router {
	return &Router{
		pingHandler,
		budgetHandler,
		teamHandler,
		job,
	}
}

func (r Router) Resource(gin *gin.Engine) {
	gin.GET("/ping", r.pingHandler.Ping)
	budget := gin.Group("/budget")
	{
		budget.GET("", r.budgetHandler.GetAll)
		budget.GET("/:code", r.budgetHandler.GetByCode)
		budget.POST("", r.budgetHandler.Create)
		budget.DELETE("/:code", r.budgetHandler.Delete)
	}

	team := gin.Group("/teams")
	{
		team.GET("/:code", r.teamHandler.GetTeamByBudget)
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

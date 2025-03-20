package team

import (
	"net/http"

	"shared-wallet-service/interfaces/dto/request"
	"shared-wallet-service/interfaces/dto/response"
	"shared-wallet-service/usecases/team"

	"github.com/gin-gonic/gin"
)

type ITeamHandler interface {
	GetTeamByBudget(c *gin.Context)
	GetTeamsByUser(c *gin.Context)
	ComposeTeam(c *gin.Context)
	DecomposeTeam(c *gin.Context)
}

type teamHandler struct {
	teamUseCase team.ITeamUseCase
}

func NewTeamHandler(teamUseCase team.ITeamUseCase) ITeamHandler {
	return &teamHandler{
		teamUseCase: teamUseCase,
	}
}

func (th *teamHandler) GetTeamByBudget(c *gin.Context) {
	ctx := c.Request.Context()
	budget, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}
	users, err := th.teamUseCase.GetTeamByBudget(ctx, budget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error searching users associated by budget",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (th *teamHandler) GetTeamsByUser(c *gin.Context) {
	ctx := c.Request.Context()
	userID, exists := c.Params.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "user's code is required",
		})
		return
	}
	teams, err := th.teamUseCase.GetTeamsByUser(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error searching budgets associated by user",
		})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (th *teamHandler) ComposeTeam(c *gin.Context) {
	ctx := c.Request.Context()
	budget, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}
	var teamRequest request.TeamRequest
	if err := c.BindJSON(&teamRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "invalid format",
		})
		return
	}
	err := th.teamUseCase.ComposeTeam(ctx, budget, teamRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error associated users with the budget",
		})
		return
	}
	c.JSON(http.StatusOK, "team composed successfully")
}

func (th *teamHandler) DecomposeTeam(c *gin.Context) {
	ctx := c.Request.Context()
	budget, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}
	var teamRequest request.TeamRequest
	if err := c.BindJSON(&teamRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "invalid format",
		})
		return
	}
	err := th.teamUseCase.DecomposeTeam(ctx, budget, teamRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error associated users with the budget",
		})
		return
	}
	c.JSON(http.StatusOK, "team decomposed successfully")
}

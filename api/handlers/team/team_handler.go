package team

import (
	"net/http"

	"cow_back/api/dto/request"
	"cow_back/api/dto/response"
	"cow_back/pkg/services/team"

	"github.com/gin-gonic/gin"
)

type ITeamHandler interface {
	GetTeamByGroup(c *gin.Context)
	GetTeamsByUser(c *gin.Context)
	ComposeTeam(c *gin.Context)
	DecomposeTeam(c *gin.Context)
}

type teamHandler struct {
	teamService team.ITeamService
}

func NewTeamHandler(teamService team.ITeamService) ITeamHandler {
	return &teamHandler{
		teamService: teamService,
	}
}

func (th *teamHandler) GetTeamByGroup(c *gin.Context) {
	ctx := c.Request.Context()
	group, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}
	users, err := th.teamService.GetTeamByGroup(ctx, group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error searching users associated by group",
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
	teams, err := th.teamService.GetTeamsByUser(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error searching groups associated by user",
		})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (th *teamHandler) ComposeTeam(c *gin.Context) {
	ctx := c.Request.Context()
	group, exists := c.Params.Get("code")
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
	err := th.teamService.ComposeTeam(ctx, group, teamRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error associated users with the group",
		})
		return
	}
	c.JSON(http.StatusOK, "team composed successfully")
}

func (th *teamHandler) DecomposeTeam(c *gin.Context) {
	ctx := c.Request.Context()
	group, exists := c.Params.Get("code")
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
	err := th.teamService.DecomposeTeam(ctx, group, teamRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error associated users with the group",
		})
		return
	}
	c.JSON(http.StatusOK, "team decomposed successfully")
}

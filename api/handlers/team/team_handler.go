package team

import (
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/pkg/services/team"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITeamHandler interface {
	GetUsersByGroup(c *gin.Context)
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

func (th *teamHandler) GetUsersByGroup(c *gin.Context) {
	ctx := c.Request.Context()
	group, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}
	users, err := th.teamService.GetUsersByGroup(ctx, group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error searching users associated by group",
		})
		return
	}
	c.JSON(http.StatusOK, users)
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

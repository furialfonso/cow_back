package handlers

import (
	"cow_back/api/dto/request"
	"cow_back/api/dto/response"
	"cow_back/pkg/services/group"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IGroupHandler interface {
	GetAll(c *gin.Context)
	GetByCode(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type groupHandler struct {
	groupService group.IGroupService
}

func NewGroupHandler(groupService group.IGroupService) IGroupHandler {
	return &groupHandler{
		groupService,
	}
}

func (gh *groupHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	groups, err := gh.groupService.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error getting groups",
		})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (gh *groupHandler) GetByCode(c *gin.Context) {
	ctx := c.Request.Context()
	code, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}

	group, err := gh.groupService.GetByCode(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error getting group %s", code),
		})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (gh *groupHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var groupRequest request.GroupRequest
	if err := c.BindJSON(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "invalid format",
		})
		return
	}
	err := gh.groupService.Create(ctx, groupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error creating group %s", groupRequest.Code),
		})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("group %s created", groupRequest.Code))
}

func (gh *groupHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	code, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}
	err := gh.groupService.Delete(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error deleting group %s", code),
		})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("group %s delete", code))
}

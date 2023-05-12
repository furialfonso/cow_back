package handlers

import (
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/pkg/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IGroupHandler interface {
	Create(c *gin.Context)
	GetGroups(c *gin.Context)
	GetGroupByCode(c *gin.Context)
}

type groupHandler struct {
	groupService services.IGroupService
}

func NewGroupHandler(groupService services.IGroupService) IGroupHandler {
	return &groupHandler{
		groupService,
	}
}

func (h *groupHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var groupDTO request.GroupDTO
	if err := c.BindJSON(&groupDTO); err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "invalid format",
		})
		return
	}
	err := h.groupService.CreateGroup(ctx, groupDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("group %s created", groupDTO.Code))
}

func (h *groupHandler) GetGroups(c *gin.Context) {
	ctx := c.Request.Context()
	groups, err := h.groupService.GetGroups(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *groupHandler) GetGroupByCode(c *gin.Context) {
	ctx := c.Request.Context()
	code, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, "grop's code is required")
		return
	}

	group, err := h.groupService.GetGroupByCode(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, group)
}

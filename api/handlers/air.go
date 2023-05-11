package handlers

import (
	"docker-go-project/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAirHandler interface {
	GetAir(c *gin.Context)
}

type airHandler struct {
	airService services.IAirService
}

func NewAirHandler(airService services.IAirService) IAirHandler {
	return &airHandler{
		airService,
	}
}

func (h *airHandler) GetAir(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := h.airService.GetAirActual(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

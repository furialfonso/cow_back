package budget

import (
	"fmt"
	"net/http"

	"shared-wallet-service/interfaces/dto/request"
	"shared-wallet-service/interfaces/dto/response"
	"shared-wallet-service/usecases/budget"

	"github.com/gin-gonic/gin"
)

type IBudgetHandler interface {
	GetAll(c *gin.Context)
	GetByCode(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type budgetHandler struct {
	budgetUseCase budget.IBudgetUseCase
}

func NewBudgetHandler(budgetUseCase budget.IBudgetUseCase) IBudgetHandler {
	return &budgetHandler{
		budgetUseCase,
	}
}

func (gh *budgetHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	budgets, err := gh.budgetUseCase.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: "error getting budgets",
		})
		return
	}
	c.JSON(http.StatusOK, budgets)
}

func (gh *budgetHandler) GetByCode(c *gin.Context) {
	ctx := c.Request.Context()
	code, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}

	budget, err := gh.budgetUseCase.GetByCode(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error getting budget %s", code),
		})
		return
	}
	c.JSON(http.StatusOK, budget)
}

func (gh *budgetHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var budgetRequest request.BudgetRequest
	if err := c.BindJSON(&budgetRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "invalid format",
		})
		return
	}
	err := gh.budgetUseCase.Create(ctx, budgetRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error creating budget %s", budgetRequest.Code),
		})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("budget %s created", budgetRequest.Code))
}

func (gh *budgetHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	code, exists := c.Params.Get("code")
	if !exists {
		c.JSON(http.StatusBadRequest, response.ApiErrors{
			Code:    http.StatusBadRequest,
			Message: "grop's code is required",
		})
		return
	}
	err := gh.budgetUseCase.Delete(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrors{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("error deleting budget %s", code),
		})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("budget %s delete", code))
}

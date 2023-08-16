package user

import (
	"docker-go-project/pkg/services/user"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetUsers(c *gin.Context)
	GetUserByCode(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type userHandler struct {
	userService user.IUserService
}

func NewUserHandler() IUserHandler {
	return &userHandler{}
}

func (uh *userHandler) GetUsers(c *gin.Context) {}

func (uh *userHandler) GetUserByCode(c *gin.Context) {}

func (uh *userHandler) Create(c *gin.Context) {}

func (uh *userHandler) Delete(c *gin.Context) {}

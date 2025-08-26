package controller

import (
	"database/sql"
	"go-chat/internal/user/repository"
	"go-chat/internal/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserServiceImpl
}

func NewUserController(db *sql.DB) *UserController {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	return &UserController{
		service: service,
	}
}

func (uc *UserController) SetUpRoutes(router *gin.Engine) {

}

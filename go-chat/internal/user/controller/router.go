package controller

import (
	"database/sql"
	"go-chat/internal/user/repository"
	"go-chat/internal/user/service"
	"go-chat/pkg/zaplog"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserServiceImpl
}

func NewUserController(db *sql.DB, log *zaplog.Zaplog) *UserController {
	repo := repository.NewUserRepository(db, log)
	service := service.NewUserService(repo)
	return &UserController{
		service: service,
	}
}

func (uc *UserController) SetUpRoutes(router *gin.Engine) {

}

package service

import (
	"go-chat/internal/user/repository"
)

type UserService interface {
}

type UserServiceImpl struct {
	repo *repository.UserRepositoryImpl
}

func NewUserService(
	repo *repository.UserRepositoryImpl,
) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

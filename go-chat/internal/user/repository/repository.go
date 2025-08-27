package repository

import (
	"database/sql"
	"go-chat/pkg/zaplog"
)

type UserRepository interface {
}

type UserRepositoryImpl struct {
	db  *sql.DB
	log *zaplog.Zaplog
}

func NewUserRepository(db *sql.DB, log *zaplog.Zaplog) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db, log: log}
}

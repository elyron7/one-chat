package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdatePasswordRequest struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type GetUserInfoRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUserInfoReponse struct {
	UUID      uuid.UUID      `json:"uuid"`      // Unique ID
	Nickname  string         `json:"nick_name"` // Nickname
	Gender    sql.NullInt64  `json:"gender"`    // Gender
	Telephone sql.NullString `json:"telephone"` // Telephone
	Email     sql.NullString `json:"email"`     // Email
	Avatar    sql.NullString `json:"avatar"`    // Avatar
	Signature sql.NullString `json:"signature"` // Signature
	Birthday  sql.NullTime   `json:"birthday"`  // Birthday
}

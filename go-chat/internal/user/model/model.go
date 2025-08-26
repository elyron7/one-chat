package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id            int       // Auto-increment primary key
	UUID          uuid.UUID // Unique ID
	Nickname      string    // Nickname
	Gender        int8      // Gender
	Telephone     string    // Phone
	Email         string    // Email
	Avatar        string    // Avatar
	Signature     string    // Signature
	Password      string    // Password
	Birthday      string    // Birthday
	CreatedAt     time.Time // Created at
	UpdatedAt     time.Time // Updated at
	LastOnlineAt  time.Time // Last online at
	LastOfflineAt time.Time // Last offline at
	Status        int8      // Status
}

package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	IsAdmin   bool            `json:"is_admin"`
	CreatedAt int64           `json:"create_at"`
	UpdatedAt int64           `json:"updated_at"`
	Details   json.RawMessage `json:"details"`
}

type Users []User

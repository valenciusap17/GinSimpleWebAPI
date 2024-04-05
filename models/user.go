package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `form:"username" json:"username,omitempty"`
	Password  string    `form:"password" json:"password,omitempty"`
	FirstName string    `form:"firstName" json:"firstName,omitempty"`
	LastName  string    `form:"lastname" json:"lastname,omitempty"`
	CreatedAt time.Time `form:"createdat" json:"createdat,omitempty"`
}

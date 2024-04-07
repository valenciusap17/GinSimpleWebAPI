package models

import (
	"time"

	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	Id       uuid.UUID         `form:"id" json:"id,omitempty"`
	UserData CreateUserRequest `form:"userdata" json:"userdata,omitempty"`
}

type CreateUserRequest struct {
	Username  string `form:"username" json:"username,omitempty"`
	Password  string `form:"password" json:"password,omitempty"`
	FirstName string `form:"firstName" json:"firstName,omitempty"`
	LastName  string `form:"lastname" json:"lastname,omitempty"`
}

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string     `form:"username" json:"username,omitempty"`
	Password  string     `form:"password" json:"password,omitempty"`
	FirstName string     `form:"firstName" json:"firstName,omitempty"`
	LastName  string     `form:"lastname" json:"lastname,omitempty"`
	CreatedAt time.Time  `form:"createdat" json:"createdat,omitempty"`
	UpdatedAt *time.Time `form:"updatedat" json:"updatedat,omitempty"`
}

func NewUser(username string, password string, firstname string, lastname string) *User {
	return &User{
		Username:  username,
		Password:  password,
		FirstName: firstname,
		LastName:  lastname,
		CreatedAt: time.Now().UTC(),
	}
}

package web

import "time"

type UserCreateRequest struct {
	LevelId   int32     `validate:"required" json:"level_id"`
	Username  string    `validate:"required,min=5,max=20" json:"username"`
	Email     string    `validate:"required" json:"email"`
	Password  string    `validate:"required" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	Id        int32     `validate:"required" json:"id"`
	LevelId   int32     `validate:"required" json:"level_id"`
	Username  string    `validate:"required,min=5,max=20" json:"username"`
	Email     string    `validate:"required" json:"email"`
	Password  string    `validate:"required" json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}

package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type UserResponse struct {
	Id        int32     `json:"id"`
	LevelId   int32     `json:"level_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		LevelId:   user.LevelId,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

package web

import "time"

type CategoryCreateRequest struct {
	Name      string    `validate:"required,min=1,max=100" json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type CategoryUpdateRequest struct {
	Id        int32     `validate:"required" json:"id" form:"id"`
	Name      string    `validate:"required,min=1,max=100" json:"name" form:"name"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

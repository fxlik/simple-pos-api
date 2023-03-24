package web

import "time"

type CategoryCreateRequest struct {
	Name      string    `validate:"required,min=1,max=100" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryUpdateRequest struct {
	Id        int32     `validate:"required" json:"id"`
	Name      string    `validate:"required,min=1,max=100" json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

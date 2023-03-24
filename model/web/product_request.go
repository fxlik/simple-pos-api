package web

import "time"

type ProductCreateRequest struct {
	CategoryId int32     `validate:"required" json:"category_id"`
	Name       string    `validate:"required, min=1,max=255" json:"name"`
	Code       string    `validate:"required, min=1,max=20" json:"code"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductUpdateRequest struct {
	Id         int32     `validate:"required" json:"id"`
	CategoryId int32     `validate:"required" json:"category_id"`
	Name       string    `validate:"required, min=1,max=255" json:"name"`
	Code       string    `validate:"required, min=1,max=20" json:"code"`
	Image      string    `json:"image"`
	UpdatedAt  time.Time `json:"updated_at"`
}

package web

import "time"

type ProductCreateRequest struct {
	CategoryId int32     `validate:"required" json:"category_id" form:"category_id"`
	Name       string    `validate:"required,min=1,max=255" json:"name" form:"name"`
	Code       string    `validate:"required,min=1,max=20" json:"code" form:"code"`
	Image      string    `json:"image" form:"image"`
	Price      int32     `validate:"required" json:"price" form:"price"`
	CreatedBy  int32     `validate:"required" json:"created_by" form:"created_by"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
}

type ProductUpdateRequest struct {
	Id         int32     `validate:"required" json:"id" form:"id"`
	CategoryId int32     `validate:"required" json:"category_id" form:"category_id"`
	Name       string    `validate:"required,min=1,max=255" json:"name" form:"name"`
	Code       string    `validate:"required,min=1,max=20" json:"code" form:"code"`
	Image      string    `json:"image" form:"image"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
}

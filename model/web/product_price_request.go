package web

import "time"

type ProductPriceCreateRequest struct {
	Price     int32     `validate:"required,min=100" json:"price" form:"price"`
	CreatedBy int32     `validate:"required" json:"created_by" form:"created_by"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type ProductPriceUpdateRequest struct {
	Id        int32     `validated:"required" json:"id" form:"id"`
	Price     int32     `validate:"required,min=100" json:"price" form:"price"`
	CreatedBy int32     `validated:"required" json:"created_by" form:"created_by"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

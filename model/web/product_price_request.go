package web

import "time"

type ProductPriceCreateRequest struct {
	Price     int32     `validate:"required,min=100" json:"price"`
	CreatedBy int32     `validate:"required" json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductPriceUpdateRequest struct {
	Id        int32     `validated:"required" json:"id"`
	Price     int32     `validate:"required,min=100" json:"price"`
	CreatedBy int32     `validated:"required" json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

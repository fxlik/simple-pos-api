package web

import "time"

type ProductStockCreateRequest struct {
	ProductId     int32     `validate:"required" json:"product_id" form:"product_id"`
	Qty           int32     `validate:"required" json:"qty" form:"qty"`
	TransactionId int32     `validate:"required" json:"transaction_id" form:"transaction_id"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
}

type ProductStockUpdateRequest struct {
	Id            int32     `validate:"required" json:"id" form:"id"`
	ProductId     int32     `validate:"required" json:"product_id" form:"product_id"`
	Qty           int32     `validate:"required" json:"qty" form:"qty"`
	TransactionId int32     `validate:"required" json:"transaction_id" form:"transaction_id"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
}

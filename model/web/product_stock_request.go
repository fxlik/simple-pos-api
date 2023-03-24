package web

import "time"

type ProductStockCreateRequest struct {
	ProductId     int32     `validate:"required" json:"product_id"`
	Qty           int32     `validate:"required" json:"qty"`
	TransactionId int32     `validate:"required" json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductStockUpdateRequest struct {
	Id            int32     `validate:"required" json:"id"`
	ProductId     int32     `validate:"required" json:"product_id"`
	Qty           int32     `validate:"required" json:"qty"`
	TransactionId int32     `validate:"required" json:"transaction_id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

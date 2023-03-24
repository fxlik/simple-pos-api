package web

import "time"

type TransactionItemCreateRequest struct {
	TransactionId int32     `validate:"required" json:"transaction_id"`
	ProductId     int32     `validate:"required" json:"product_id"`
	Qty           int32     `validate:"required" json:"qty"`
	Price         int32     `json:"price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TransactionItemCreateResponse struct {
	Id            int32     `validate:"required" json:"id"`
	TransactionId int32     `validate:"required" json:"transaction_id"`
	ProductId     int32     `validate:"required" json:"product_id"`
	Qty           int32     `validate:"required" json:"qty"`
	Price         int32     `json:"price"`
	UpdatedAt     time.Time `json:"updated_at"`
}

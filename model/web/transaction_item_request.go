package web

import "time"

type TransactionItemCreateRequest struct {
	TransactionId int32 `validate:"required" json:"transaction_id" form:"transaction_id"`
	ProductId     int32 `validate:"required" json:"product_id" form:"product_id"`
	Qty           int32 `validate:"required" json:"qty" form:"qty"`
	//Price         int32     `json:"price" form:"price"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type TransactionItemUpdateRequest struct {
	Id            int32     `validate:"required" json:"id" form:"id"`
	TransactionId int32     `validate:"required" json:"transaction_id" form:"transaction_id"`
	ProductId     int32     `validate:"required" json:"product_id" form:"product_id"`
	Qty           int32     `validate:"required" json:"qty" form:"qty"`
	Price         int32     `json:"price" form:"price"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
}

package web

import "time"

type InvoiceCreateRequest struct {
	SubTotal      int32     `validate:"required" json:"sub_total" form:"sub_total"`
	TransactionId int32     `validate:"required" json:"transaction_id" form:"transaction_id"`
	CreatedBy     int32     `validate:"required" json:"created_by" form:"created_by"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
}

type InvoiceUpdateRequest struct {
	Id            int32     `validate:"required" json:"id" form:"id"`
	SubTotal      int32     `validate:"required" json:"sub_total" form:"sub_total"`
	TransactionId int32     `validate:"required" json:"transaction_id" form:"transaction_id"`
	CreatedBy     int32     `validate:"required" json:"created_by" form:"created_by"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
}

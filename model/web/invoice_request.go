package web

import "time"

type InvoiceCreateRequest struct {
	SubTotal      int32     `validate:"required" json:"sub_total"`
	TransactionId int32     `validate:"required" json:"transaction_id"`
	CreatedBy     int32     `validate:"required" json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type InvoiceUpdateRequest struct {
	Id            int32     `validate:"required" json:"id"`
	SubTotal      int32     `validate:"required" json:"sub_total"`
	TransactionId int32     `validate:"required" json:"transaction_id"`
	CreatedBy     int32     `validate:"required" json:"created_by"`
	UpdatedAt     time.Time `json:"updated_at"`
}

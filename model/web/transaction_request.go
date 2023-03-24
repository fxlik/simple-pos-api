package web

import "time"

type TransactionCreateRequest struct {
	Code                string    `validate:"required" json:"code"`
	TransactionTypeId   int32     `validate:"required" json:"transaction_type_id"`
	SupplierId          int32     `validate:"required" json:"supplier_id"`
	TransactionStatusId int32     `validate:"required" json:"transaction_status_id"`
	CreatedBy           int32     `validate:"required" json:"created_by"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type TransactionUpdateRequest struct {
	Id                  int32     `json:"id"`
	Code                string    `validate:"required" json:"code"`
	TransactionTypeId   int32     `validate:"required" json:"transaction_type_id"`
	SupplierId          int32     `validate:"required" json:"supplier_id"`
	TransactionStatusId int32     `validate:"required" json:"transaction_status_id"`
	CreatedBy           int32     `validate:"required" json:"created_by"`
	UpdatedAt           time.Time `json:"updated_at"`
}
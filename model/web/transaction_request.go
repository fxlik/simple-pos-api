package web

import (
	"time"
)

type TransactionCreateRequest struct {
	Code                string    `json:"code" form:"code"`
	TransactionTypeId   int32     `validate:"required" json:"transaction_type_id" form:"transaction_type_id"`
	SupplierId          *int32    `json:"supplier_id" form:"supplier_id"`
	TransactionStatusId int32     `validate:"required" json:"transaction_status_id" form:"transaction_status_id"`
	CreatedBy           int32     `validate:"required" json:"created_by" form:"created_by"`
	CreatedAt           time.Time `json:"created_at" form:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" form:"updated_at"`
}

type TransactionUpdateRequest struct {
	Id                  int32     `json:"id" form:"id"`
	Code                string    `json:"code" form:"code"`
	TransactionTypeId   int32     `validate:"required" json:"transaction_type_id" form:"transaction_type_id"`
	SupplierId          *int32    `json:"supplier_id" form:"supplier_id"`
	TransactionStatusId int32     `validate:"required" json:"transaction_status_id" form:"transaction_status_id"`
	CreatedBy           int32     `validate:"required" json:"created_by" form:"created_by"`
	UpdatedAt           time.Time `json:"updated_at" form:"updated_at"`
}

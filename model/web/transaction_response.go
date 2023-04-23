package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type TransactionResponse struct {
	Id                  int32     `json:"id"`
	Code                string    `json:"code"`
	TransactionTypeId   int32     `json:"transaction_type_id"`
	SupplierId          *int32    `json:"supplier_id"`
	TransactionStatusId int32     `json:"transaction_status_id"`
	CreatedBy           int32     `json:"created_by"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func ToTransactionResponse(transaction domain.Transaction) TransactionResponse {
	return TransactionResponse{
		Id:                  transaction.Id,
		Code:                transaction.Code,
		TransactionTypeId:   transaction.TransactionTypeId,
		SupplierId:          transaction.SupplierId,
		TransactionStatusId: transaction.TransactionStatusId,
		CreatedBy:           transaction.CreatedBy,
		CreatedAt:           transaction.CreatedAt,
		UpdatedAt:           transaction.UpdatedAt,
	}
}

func ToTransactionResponses(transactions []domain.Transaction) []TransactionResponse {
	var transactionResponses []TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, ToTransactionResponse(transaction))
	}
	return transactionResponses
}

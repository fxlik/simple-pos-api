package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type TransactionItemResponse struct {
	Id            int32     `json:"id"`
	TransactionId int32     `json:"transaction_id"`
	ProductId     int32     `json:"product_id"`
	Qty           int32     `json:"qty"`
	Price         int32     `json:"price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToTransactionItemResponse(item domain.TransactionItem) TransactionItemResponse {
	return TransactionItemResponse{
		Id:            item.Id,
		TransactionId: item.TransactionId,
		ProductId:     item.ProductId,
		Qty:           item.Qty,
		Price:         item.Price,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}
}

func ToTransactionItemResponses(items []domain.TransactionItem) []TransactionItemResponse {
	var transactionItemResponses []TransactionItemResponse
	for _, item := range items {
		transactionItemResponses = append(transactionItemResponses, ToTransactionItemResponse(item))
	}
	return transactionItemResponses
}

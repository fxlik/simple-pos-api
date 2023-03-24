package web

import "fxlik/simple-post-api/model/domain"

type TransactionTypeResponse struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Disabled bool   `json:"disabled"`
}

func ToTransactionTypeResponse(transactionType domain.TransactionType) TransactionTypeResponse {
	return TransactionTypeResponse{
		Id:       transactionType.Id,
		Name:     transactionType.Name,
		Disabled: transactionType.Disabled,
	}
}

func ToTransactionTypeResponses(transactionTypes []domain.TransactionType) []TransactionTypeResponse {
	var transactionTypeResponses []TransactionTypeResponse
	for _, transactionType := range transactionTypes {
		transactionTypeResponses = append(transactionTypeResponses, ToTransactionTypeResponse(transactionType))
	}
	return transactionTypeResponses
}

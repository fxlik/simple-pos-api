package web

import "fxlik/simple-post-api/model/domain"

type TransactionStatusResponse struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Disabled bool   `json:"disabled"`
}

func ToTransactionStatusResponse(status domain.TransactionStatus) TransactionStatusResponse {
	return TransactionStatusResponse{
		Id:       status.Id,
		Name:     status.Name,
		Disabled: status.Disabled,
	}
}

func ToTransactionStatusResponses(statuses []domain.TransactionStatus) []TransactionStatusResponse {
	var transactionStatusResponses []TransactionStatusResponse
	for _, status := range statuses {
		transactionStatusResponses = append(transactionStatusResponses, ToTransactionStatusResponse(status))
	}
	return transactionStatusResponses
}

package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type InvoiceResponse struct {
	Id            int32     `json:"id"`
	SubTotal      int32     `json:"sub_total"`
	TransactionId int32     `json:"transaction_id"`
	CreatedBy     int32     `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToInvoiceResponse(invoice domain.Invoice) InvoiceResponse {
	return InvoiceResponse{
		Id:            invoice.Id,
		SubTotal:      invoice.SubTotal,
		TransactionId: invoice.TransactionId,
		CreatedBy:     invoice.CreatedBy,
		CreatedAt:     invoice.CreatedAt,
		UpdatedAt:     invoice.UpdatedAt,
	}
}

func ToInvoiceResponses(invoices []domain.Invoice) []InvoiceResponse {
	var invoiceResponses []InvoiceResponse
	for _, invoice := range invoices {
		invoiceResponses = append(invoiceResponses, ToInvoiceResponse(invoice))
	}
	return invoiceResponses
}

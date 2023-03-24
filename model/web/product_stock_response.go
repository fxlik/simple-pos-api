package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type ProductStockResponse struct {
	Id            int32     `json:"id"`
	ProductId     int32     `json:"product_id"`
	Qty           int32     `json:"qty"`
	TransactionId int32     `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToProductStockResponse(stock domain.ProductStock) ProductStockResponse {
	return ProductStockResponse{
		Id:            stock.Id,
		ProductId:     stock.ProductId,
		Qty:           stock.Qty,
		TransactionId: stock.TransactionId,
		CreatedAt:     stock.CreatedAt,
		UpdatedAt:     stock.UpdatedAt,
	}
}

func ToProductStockResponses(stocks []domain.ProductStock) []ProductStockResponse {
	var productStockResponses []ProductStockResponse
	for _, stock := range stocks {
		productStockResponses = append(productStockResponses, ToProductStockResponse(stock))
	}
	return productStockResponses
}

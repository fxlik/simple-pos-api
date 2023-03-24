package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type ProductPriceResponse struct {
	Id        int32     `json:"id"`
	Price     int32     `json:"price"`
	CreatedBy int32     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToProductPriceResponse(productPrice domain.ProductPrice) ProductPriceResponse {
	return ProductPriceResponse{
		Id:        productPrice.Id,
		Price:     productPrice.Price,
		CreatedBy: productPrice.CreatedBy,
		CreatedAt: productPrice.CreatedAt,
		UpdatedAt: productPrice.UpdatedAt,
	}
}

func ToProductPriceResponses(productPrices []domain.ProductPrice) []ProductPriceResponse {
	var productPriceResponses []ProductPriceResponse
	for _, productPrice := range productPrices {
		productPriceResponses = append(productPriceResponses, ToProductPriceResponse(productPrice))
	}
	return productPriceResponses
}

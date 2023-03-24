package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type ProductResponse struct {
	Id         int32     `json:"id"`
	CategoryId int32     `json:"category_id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Id:         product.Id,
		CategoryId: product.CategoryId,
		Name:       product.Name,
		Code:       product.Code,
		Image:      product.Image,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func ToProductResponses(products []domain.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

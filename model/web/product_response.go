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

type ProductWithPriceResponse struct {
	Id           int32                 `json:"id"`
	CategoryId   int32                 `json:"category_id"`
	Name         string                `json:"name"`
	Code         string                `json:"code"`
	Image        string                `json:"image"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	ProductPrice *ProductPriceResponse `json:"product_price"`
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

func ToProductWithPriceResponse(product domain.Product, productPrice domain.ProductPrice) ProductWithPriceResponse {
	var p ProductWithPriceResponse
	p.ProductPrice = nil
	if productPrice.Id != 0 {
		p.ProductPrice = &ProductPriceResponse{
			Id:        productPrice.Id,
			ProductId: productPrice.ProductId,
			Price:     productPrice.Price,
			CreatedBy: productPrice.CreatedBy,
			CreatedAt: productPrice.CreatedAt,
			UpdatedAt: productPrice.UpdatedAt,
		}
	}
	return ProductWithPriceResponse{
		Id:           product.Id,
		CategoryId:   product.CategoryId,
		Name:         product.Name,
		Code:         product.Code,
		Image:        product.Image,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
		ProductPrice: p.ProductPrice,
	}
}

func ToProductResponses(products []domain.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

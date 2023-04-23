package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type ProductService interface {
	Save(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productId int32)
	FindById(ctx context.Context, productId int32) web.ProductWithPriceResponse
	FindAll(ctx context.Context) []web.ProductResponse
}

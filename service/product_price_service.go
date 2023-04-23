package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type ProductPriceService interface {
	Save(ctx context.Context, request web.ProductPriceCreateRequest) web.ProductPriceResponse
	Update(ctx context.Context, request web.ProductPriceUpdateRequest) web.ProductPriceResponse
	Delete(ctx context.Context, productPriceId int32)
	FindById(ctx context.Context, productPriceId int32) web.ProductPriceResponse
	FindOneByProductId(ctx context.Context, productId int32) web.ProductPriceResponse
	FindAll(ctx context.Context) []web.ProductPriceResponse
}

package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type SupplierService interface {
	Save(ctx context.Context, request web.SupplierCreateRequest) web.SupplierResponse
	Update(ctx context.Context, request web.SupplierUpdateRequest) web.SupplierResponse
	Delete(ctx context.Context, supplierId int32)
	FindById(ctx context.Context, supplierId int32) web.SupplierResponse
	FindAll(ctx context.Context) []web.SupplierResponse
}

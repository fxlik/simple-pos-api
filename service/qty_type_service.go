package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type QtyTypeService interface {
	Save(ctx context.Context, request web.QtyTypeCreateRequest) web.QtyTypeResponse
	Update(ctx context.Context, request web.QtyTypeUpdateRequest) web.QtyTypeResponse
	Delete(ctx context.Context, qtyTypeId int32)
	FindById(ctx context.Context, qtyTypeId int32) web.QtyTypeResponse
	FindAll(ctx context.Context) []web.QtyTypeResponse
}

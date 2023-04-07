package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type TransactionTypeService interface {
	Save(ctx context.Context, request web.TransactionTypeCreateRequest) web.TransactionTypeResponse
	Update(ctx context.Context, request web.TransactionTypeUpdateRequest) web.TransactionTypeResponse
	Delete(ctx context.Context, transactionTypeId int32)
	FindById(ctx context.Context, transactionTypeId int32) web.TransactionTypeResponse
	FindAll(ctx context.Context) []web.TransactionTypeResponse
}

package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type TransactionItemService interface {
	Save(ctx context.Context, request web.TransactionItemCreateRequest) web.TransactionItemResponse
	Update(ctx context.Context, request web.TransactionItemUpdateRequest) web.TransactionItemResponse
	Delete(ctx context.Context, transactionItemId int32)
	FindById(ctx context.Context, transactionItemId int32) web.TransactionItemResponse
	FindAll(ctx context.Context) []web.TransactionItemResponse
}

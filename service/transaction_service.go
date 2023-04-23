package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type TransactionService interface {
	Save(ctx context.Context, request web.TransactionCreateRequest) web.TransactionResponse
	Update(ctx context.Context, request web.TransactionUpdateRequest) web.TransactionResponse
	Delete(ctx context.Context, transactionId int32)
	FindById(ctx context.Context, transactionId int32) web.TransactionResponse
	FindAll(ctx context.Context) []web.TransactionResponse
}

package service

import (
	"context"
	"fxlik/simple-post-api/model/web"
)

type TransactionStatusService interface {
	Save(ctx context.Context, request web.TransactionStatusCreateRequest) web.TransactionStatusResponse
	Update(ctx context.Context, request web.TransactionStatusUpdateRequest) web.TransactionStatusResponse
	Delete(ctx context.Context, transactionStatusId int32)
	FindById(ctx context.Context, transactionStatusId int32) web.TransactionStatusResponse
	FindAll(ctx context.Context) []web.TransactionStatusResponse
}

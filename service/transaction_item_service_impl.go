package service

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/exception"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/repository"
	"github.com/go-playground/validator/v10"
)

type TransactionItemServiceImpl struct {
	TransactionItemRepository repository.TransactionItemRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewTransactionItemServiceImpl(transactionItemRepository repository.TransactionItemRepository, DB *sql.DB, validate *validator.Validate) *TransactionItemServiceImpl {
	return &TransactionItemServiceImpl{
		TransactionItemRepository: transactionItemRepository,
		DB:                        DB,
		Validate:                  validate,
	}
}

func (service *TransactionItemServiceImpl) Save(ctx context.Context, request web.TransactionItemCreateRequest) web.TransactionItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	transactionItem := domain.TransactionItem{
		TransactionId: request.TransactionId,
		ProductId:     request.ProductId,
		Qty:           request.Qty,
		Price:         request.Price,
		CreatedAt:     helper.GetTime(),
		UpdatedAt:     helper.GetTime(),
	}
	transactionItem = service.TransactionItemRepository.Save(ctx, tx, transactionItem)
	return web.ToTransactionItemResponse(transactionItem)
}

func (service *TransactionItemServiceImpl) Update(ctx context.Context, request web.TransactionItemUpdateRequest) web.TransactionItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	transactionItem, errFind := service.TransactionItemRepository.FindById(ctx, tx, request.Id)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}
	transactionItem.TransactionId = request.TransactionId
	transactionItem.ProductId = request.ProductId
	transactionItem.Qty = request.Qty
	transactionItem.Price = request.Price
	transactionItem.UpdatedAt = helper.GetTime()

	transactionItem = service.TransactionItemRepository.Update(ctx, tx, transactionItem)
	return web.ToTransactionItemResponse(transactionItem)
}

func (service *TransactionItemServiceImpl) Delete(ctx context.Context, transactionItemId int32) {
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	transactionItem, errFind := service.TransactionItemRepository.FindById(ctx, tx, transactionItemId)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}
	service.TransactionItemRepository.Delete(ctx, tx, transactionItem)
}

func (service *TransactionItemServiceImpl) FindById(ctx context.Context, transactionItemId int32) web.TransactionItemResponse {
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	transactionItem, errFind := service.TransactionItemRepository.FindById(ctx, tx, transactionItemId)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}
	return web.ToTransactionItemResponse(transactionItem)
}

func (service *TransactionItemServiceImpl) FindAll(ctx context.Context) []web.TransactionItemResponse {
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	transactionItems := service.TransactionItemRepository.FindAll(ctx, tx)
	return web.ToTransactionItemResponses(transactionItems)
}

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

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewTransactionServiceImpl(transactionRepository repository.TransactionRepository, DB *sql.DB, validate *validator.Validate) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) Save(ctx context.Context, request web.TransactionCreateRequest) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction := domain.Transaction{
		Code:                helper.GenCode(request.TransactionTypeId),
		TransactionTypeId:   request.TransactionTypeId,
		SupplierId:          request.SupplierId,
		TransactionStatusId: request.TransactionStatusId,
		CreatedBy:           request.CreatedBy,
		CreatedAt:           helper.GetTime(),
		UpdatedAt:           helper.GetTime(),
	}
	transaction = service.TransactionRepository.Save(ctx, tx, transaction)
	return web.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Update(ctx context.Context, request web.TransactionUpdateRequest) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	transaction.TransactionTypeId = request.TransactionTypeId
	transaction.SupplierId = request.SupplierId
	transaction.TransactionStatusId = request.TransactionStatusId
	transaction.CreatedBy = request.CreatedBy
	transaction.UpdatedAt = helper.GetTime()
	transaction = service.TransactionRepository.Update(ctx, tx, transaction)
	return web.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Delete(ctx context.Context, transactionId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.TransactionRepository.Delete(ctx, tx, transaction)
}

func (service *TransactionServiceImpl) FindById(ctx context.Context, transactionId int32) web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindAll(ctx context.Context) []web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactions := service.TransactionRepository.FindAll(ctx, tx)
	return web.ToTransactionResponses(transactions)
}

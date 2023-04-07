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

type TransactionStatusServiceImpl struct {
	TransactionStatusRepository repository.TransactionStatusRepository
	DB                          *sql.DB
	Validate                    *validator.Validate
}

func NewTransactionStatusServiceImpl(transactionStatusRepository repository.TransactionStatusRepository, DB *sql.DB, validate *validator.Validate) *TransactionStatusServiceImpl {
	return &TransactionStatusServiceImpl{
		TransactionStatusRepository: transactionStatusRepository,
		DB:                          DB,
		Validate:                    validate,
	}
}

func (service *TransactionStatusServiceImpl) Save(ctx context.Context, request web.TransactionStatusCreateRequest) web.TransactionStatusResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionStatus := domain.TransactionStatus{
		Name:     request.Name,
		Disabled: request.Disabled,
	}

	transactionStatus = service.TransactionStatusRepository.Save(ctx, tx, transactionStatus)
	return web.ToTransactionStatusResponse(transactionStatus)
}

func (service *TransactionStatusServiceImpl) Update(ctx context.Context, request web.TransactionStatusUpdateRequest) web.TransactionStatusResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionStatus, err := service.TransactionStatusRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	transactionStatus.Name = request.Name
	transactionStatus.Disabled = request.Disabled
	transactionStatus = service.TransactionStatusRepository.Update(ctx, tx, transactionStatus)
	return web.ToTransactionStatusResponse(transactionStatus)
}

func (service *TransactionStatusServiceImpl) Delete(ctx context.Context, transactionStatusId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionStatus, err := service.TransactionStatusRepository.FindById(ctx, tx, transactionStatusId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.TransactionStatusRepository.Delete(ctx, tx, transactionStatus)
}

func (service *TransactionStatusServiceImpl) FindById(ctx context.Context, transactionStatusId int32) web.TransactionStatusResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	transactionStatus, err := service.TransactionStatusRepository.FindById(ctx, tx, transactionStatusId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToTransactionStatusResponse(transactionStatus)
}

func (service *TransactionStatusServiceImpl) FindAll(ctx context.Context) []web.TransactionStatusResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionStatuses := service.TransactionStatusRepository.FindAll(ctx, tx)
	return web.ToTransactionStatusResponses(transactionStatuses)
}

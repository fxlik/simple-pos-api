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

type TransactionTypeServiceImpl struct {
	TransactionTypeRepository repository.TransactionTypeRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewTransactionTypeServiceImpl(transactionTypeRepository repository.TransactionTypeRepository, DB *sql.DB, validate *validator.Validate) *TransactionTypeServiceImpl {
	return &TransactionTypeServiceImpl{
		TransactionTypeRepository: transactionTypeRepository,
		DB:                        DB,
		Validate:                  validate,
	}
}

func (service *TransactionTypeServiceImpl) Save(ctx context.Context, request web.TransactionTypeCreateRequest) web.TransactionTypeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionType := domain.TransactionType{
		Name:     request.Name,
		Disabled: request.Disabled,
	}

	transactionType = service.TransactionTypeRepository.Save(ctx, tx, transactionType)
	return web.ToTransactionTypeResponse(transactionType)
}

func (service *TransactionTypeServiceImpl) Update(ctx context.Context, request web.TransactionTypeUpdateRequest) web.TransactionTypeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionType, err := service.TransactionTypeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	transactionType.Name = request.Name
	transactionType.Disabled = request.Disabled
	transactionType = service.TransactionTypeRepository.Update(ctx, tx, transactionType)
	return web.ToTransactionTypeResponse(transactionType)
}

func (service *TransactionTypeServiceImpl) Delete(ctx context.Context, transactionTypeId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionType, err := service.TransactionTypeRepository.FindById(ctx, tx, transactionTypeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.TransactionTypeRepository.Delete(ctx, tx, transactionType)
}

func (service *TransactionTypeServiceImpl) FindById(ctx context.Context, transactionTypeId int32) web.TransactionTypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionType, err := service.TransactionTypeRepository.FindById(ctx, tx, transactionTypeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToTransactionTypeResponse(transactionType)
}

func (service *TransactionTypeServiceImpl) FindAll(ctx context.Context) []web.TransactionTypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactionTypes := service.TransactionTypeRepository.FindAll(ctx, tx)
	return web.ToTransactionTypeResponses(transactionTypes)
}

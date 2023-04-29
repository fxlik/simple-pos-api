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
	TransactionRepository     repository.TransactionRepository
	ProductPriceRepository    repository.ProductPriceRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewTransactionItemServiceImpl(transactionItemRepository repository.TransactionItemRepository, transactionRepository repository.TransactionRepository, productPriceRepository repository.ProductPriceRepository, DB *sql.DB, validate *validator.Validate) *TransactionItemServiceImpl {
	return &TransactionItemServiceImpl{
		TransactionItemRepository: transactionItemRepository,
		TransactionRepository:     transactionRepository,
		ProductPriceRepository:    productPriceRepository,
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

	//FIND PRODUCT PRICE FIRST
	productPrice, errFindProductPrice := service.ProductPriceRepository.FindOneByProductId(ctx, tx, request.ProductId)
	if errFindProductPrice != nil {
		panic(exception.NewNotFoundError(errFindProductPrice.Error()))
	}

	//FIND IF TRANSACTION ITEM IS EXIST WITH THE SAME TRANSACTION AND PRODUCT ID
	//RETURN IS BOOLEAN (TRUE OR FALSE)
	transactionItemCheck, errCheck := service.TransactionItemRepository.CheckIfExistByTransactionIdAndProductId(ctx, tx, request.TransactionId, request.ProductId)
	if errCheck != nil {
		helper.PanicIfError(errCheck)
	}

	transactionItem := domain.TransactionItem{}

	if transactionItemCheck == true {
		transactionItemExist, _ := service.TransactionItemRepository.FindOneByTransactionIdAndProductId(ctx, tx, request.TransactionId, request.ProductId)
		transactionItem.Id = transactionItemExist.Id
		transactionItem.TransactionId = request.TransactionId
		transactionItem.ProductId = request.ProductId
		transactionItem.Qty = transactionItemExist.Qty + request.Qty
		transactionItem.Price = productPrice.Price
		transactionItem.UpdatedAt = helper.GetTime()

		transactionItem = service.TransactionItemRepository.Update(ctx, tx, transactionItem)
	} else {
		//MAKE A NEW ITEM
		transactionItem.TransactionId = request.TransactionId
		transactionItem.ProductId = request.ProductId
		transactionItem.Qty = request.Qty
		transactionItem.Price = productPrice.Price
		transactionItem.CreatedAt = helper.GetTime()
		transactionItem.UpdatedAt = helper.GetTime()

		transactionItem = service.TransactionItemRepository.Save(ctx, tx, transactionItem)
	}

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

func (service *TransactionItemServiceImpl) FindAllByTransactionId(ctx context.Context, transactionId int32) []web.TransactionItemResponse {
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	//check if transaction is exists
	_, errFindTransaction := service.TransactionRepository.FindById(ctx, tx, transactionId)
	if errFindTransaction != nil {
		panic(exception.NewNotFoundError(errFindTransaction.Error()))
	}

	transactionItems := service.TransactionItemRepository.FindAllByTransactionId(ctx, tx, transactionId)
	return web.ToTransactionItemResponses(transactionItems)
}

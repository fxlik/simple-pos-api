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

type ProductPriceServiceImpl struct {
	ProductPriceRepository repository.ProductPriceRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewProductPriceServiceImpl(productPriceRepository repository.ProductPriceRepository, DB *sql.DB, validate *validator.Validate) *ProductPriceServiceImpl {
	return &ProductPriceServiceImpl{
		ProductPriceRepository: productPriceRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *ProductPriceServiceImpl) Save(ctx context.Context, request web.ProductPriceCreateRequest) web.ProductPriceResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productPrice := domain.ProductPrice{
		ProductId: request.ProductId,
		Price:     request.Price,
		CreatedBy: request.CreatedBy,
		CreatedAt: helper.GetTime(),
		UpdatedAt: helper.GetTime(),
	}

	productPrice = service.ProductPriceRepository.Save(ctx, tx, productPrice)
	return web.ToProductPriceResponse(productPrice)
}

func (service *ProductPriceServiceImpl) Update(ctx context.Context, request web.ProductPriceUpdateRequest) web.ProductPriceResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productPrice, err := service.ProductPriceRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	productPrice.ProductId = request.ProductId
	productPrice.Price = request.Price
	productPrice.CreatedBy = request.CreatedBy
	productPrice.UpdatedAt = request.UpdatedAt
	productPrice = service.ProductPriceRepository.Update(ctx, tx, productPrice)
	return web.ToProductPriceResponse(productPrice)
}

func (service *ProductPriceServiceImpl) Delete(ctx context.Context, productPriceId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productPrice, err := service.ProductPriceRepository.FindById(ctx, tx, productPriceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.ProductPriceRepository.Delete(ctx, tx, productPrice)
}

func (service *ProductPriceServiceImpl) FindById(ctx context.Context, productPriceId int32) web.ProductPriceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productPrice, err := service.ProductPriceRepository.FindById(ctx, tx, productPriceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToProductPriceResponse(productPrice)
}

func (service *ProductPriceServiceImpl) FindOneByProductId(ctx context.Context, productId int32) web.ProductPriceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productPrice, err := service.ProductPriceRepository.FindOneByProductId(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToProductPriceResponse(productPrice)
}

func (service *ProductPriceServiceImpl) FindAll(ctx context.Context) []web.ProductPriceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productPrices := service.ProductPriceRepository.FindAll(ctx, tx)
	return web.ToProductPriceResponses(productPrices)
}

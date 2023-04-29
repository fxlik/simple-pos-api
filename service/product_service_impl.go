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

type ProductServiceImpl struct {
	ProductRepository      repository.ProductRepository
	ProductPriceRepository repository.ProductPriceRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, productPriceRepository repository.ProductPriceRepository, DB *sql.DB, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository:      productRepository,
		ProductPriceRepository: productPriceRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *ProductServiceImpl) Save(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	errValidate := service.Validate.Struct(request)
	//helper.PanicIfError(errValidate)
	if errValidate != nil {
		panic(errValidate.Error())
		helper.DeleteFile(request.Image)
	}
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		CategoryId: request.CategoryId,
		Name:       request.Name,
		Code:       request.Code,
		Image:      request.Image,
		CreatedAt:  helper.GetTime(),
		UpdatedAt:  helper.GetTime(),
	}

	product = service.ProductRepository.Save(ctx, tx, product)
	productPrice := domain.ProductPrice{
		ProductId: product.Id,
		Price:     request.Price,
		CreatedBy: request.CreatedBy,
		CreatedAt: helper.GetTime(),
		UpdatedAt: helper.GetTime(),
	}
	productPrice = service.ProductPriceRepository.Save(ctx, tx, productPrice)
	return web.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		//delete uploaded file if product is not found
		helper.DeleteFile(request.Image)
		panic(exception.NewNotFoundError(err.Error()))
	}

	errFile := helper.DeleteFile(product.Image)
	if errFile != nil {
		helper.PanicIfError(errFile)
	}

	product.CategoryId = request.CategoryId
	product.Name = request.Name
	product.Code = request.Code
	product.Image = request.Image
	product.UpdatedAt = helper.GetTime()

	product = service.ProductRepository.Update(ctx, tx, product)
	return web.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int32) {
	tx, errDb := service.DB.Begin()
	helper.PanicIfError(errDb)
	defer helper.CommitOrRollback(tx)

	product, errFind := service.ProductRepository.FindById(ctx, tx, productId)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}
	errFile := helper.DeleteFile(product.Image)
	if errFile != nil {
		helper.PanicIfError(errFile)
	}
	service.ProductPriceRepository.DeleteByProductId(ctx, tx, productId)
	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int32) web.ProductWithPriceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	productPrice, _ := service.ProductPriceRepository.FindOneByProductId(ctx, tx, productId)

	return web.ToProductWithPriceResponse(product, productPrice)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)
	return web.ToProductResponses(products)
}

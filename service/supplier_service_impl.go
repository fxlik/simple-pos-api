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

type SupplierServiceImpl struct {
	SupplierRepository repository.SupplierRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewSupplierServiceImpl(supplierRepository repository.SupplierRepository, DB *sql.DB, validator *validator.Validate) *SupplierServiceImpl {
	return &SupplierServiceImpl{
		SupplierRepository: supplierRepository,
		DB:                 DB,
		Validator:          validator,
	}
}

func (service *SupplierServiceImpl) Save(ctx context.Context, request web.SupplierCreateRequest) web.SupplierResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier := domain.Supplier{
		Name:      request.Name,
		Email:     request.Email,
		Telp:      request.Telp,
		CreatedAt: helper.GetTime(),
		UpdatedAt: helper.GetTime(),
	}
	supplier = service.SupplierRepository.Save(ctx, tx, supplier)
	return web.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) Update(ctx context.Context, request web.SupplierUpdateRequest) web.SupplierResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier, err := service.SupplierRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	supplier.Name = request.Name
	supplier.Email = request.Email
	supplier.Telp = request.Telp
	supplier.UpdatedAt = helper.GetTime()
	supplier = service.SupplierRepository.Update(ctx, tx, supplier)
	return web.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) Delete(ctx context.Context, supplierId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier, err := service.SupplierRepository.FindById(ctx, tx, supplierId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.SupplierRepository.Delete(ctx, tx, supplier)
}

func (service *SupplierServiceImpl) FindById(ctx context.Context, supplierId int32) web.SupplierResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier, err := service.SupplierRepository.FindById(ctx, tx, supplierId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) FindAll(ctx context.Context) []web.SupplierResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	suppliers := service.SupplierRepository.FindAll(ctx, tx)
	return web.ToSupplierResponses(suppliers)
}

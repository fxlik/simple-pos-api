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

type QtyTypeServiceImpl struct {
	QtyTypeRepository repository.QtyTypeRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewQtyTypeServiceImpl(qtyTypeRepository repository.QtyTypeRepository, DB *sql.DB, validate *validator.Validate) *QtyTypeServiceImpl {
	return &QtyTypeServiceImpl{
		QtyTypeRepository: qtyTypeRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *QtyTypeServiceImpl) Save(ctx context.Context, request web.QtyTypeCreateRequest) web.QtyTypeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	qtyType := domain.QtyType{
		Name:      request.Name,
		Disabled:  request.Disabled,
		CreatedAt: helper.GetTime(),
		UpdatedAt: helper.GetTime(),
	}

	qtyType = service.QtyTypeRepository.Save(ctx, tx, qtyType)
	return web.ToQtyTypeResponse(qtyType)
}

func (service *QtyTypeServiceImpl) Update(ctx context.Context, request web.QtyTypeUpdateRequest) web.QtyTypeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	qtyType, err := service.QtyTypeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	qtyType.Name = request.Name
	qtyType.Disabled = request.Disabled
	qtyType.UpdatedAt = helper.GetTime()

	qtyType = service.QtyTypeRepository.Update(ctx, tx, qtyType)
	return web.ToQtyTypeResponse(qtyType)
}

func (service *QtyTypeServiceImpl) Delete(ctx context.Context, qtyTypeId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	qtyType, err := service.QtyTypeRepository.FindById(ctx, tx, qtyTypeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.QtyTypeRepository.Delete(ctx, tx, qtyType)
}

func (service *QtyTypeServiceImpl) FindById(ctx context.Context, qtyTypeId int32) web.QtyTypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	qtyType, err := service.QtyTypeRepository.FindById(ctx, tx, qtyTypeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToQtyTypeResponse(qtyType)
}

func (service *QtyTypeServiceImpl) FindAll(ctx context.Context) []web.QtyTypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	qtyTypes := service.QtyTypeRepository.FindAll(ctx, tx)
	return web.ToQtyTypeResponses(qtyTypes)
}

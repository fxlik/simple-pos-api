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

type LevelServiceImpl struct {
	LevelRepository repository.LevelRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewLevelServiceImpl(levelRepository repository.LevelRepository, DB *sql.DB, validate *validator.Validate) *LevelServiceImpl {
	return &LevelServiceImpl{
		LevelRepository: levelRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *LevelServiceImpl) Save(ctx context.Context, request web.LevelCreateRequest) web.LevelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	level := domain.Level{
		Name: request.Name,
	}
	level = service.LevelRepository.Save(ctx, tx, level)
	return web.ToLevelResponse(level)
}

func (service *LevelServiceImpl) Update(ctx context.Context, request web.LevelUpdateRequest) web.LevelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	level, err := service.LevelRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	level.Name = request.Name
	level = service.LevelRepository.Update(ctx, tx, level)
	return web.ToLevelResponse(level)
}

func (service *LevelServiceImpl) Delete(ctx context.Context, levelId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	level, err := service.LevelRepository.FindById(ctx, tx, levelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.LevelRepository.Delete(ctx, tx, level)
}

func (service *LevelServiceImpl) FindById(ctx context.Context, levelId int32) web.LevelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	level, err := service.LevelRepository.FindById(ctx, tx, levelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.ToLevelResponse(level)
}

func (service *LevelServiceImpl) FindAll(ctx context.Context) []web.LevelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	levels := service.LevelRepository.FindAll(ctx, tx)
	return web.ToLevelResponses(levels)
}

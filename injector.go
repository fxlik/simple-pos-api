//go:build wireinject
// +build wireinject

package main

import (
	"fxlik/simple-post-api/app"
	"fxlik/simple-post-api/controller"
	"fxlik/simple-post-api/repository"
	"fxlik/simple-post-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

var levelSet = wire.NewSet(
	repository.NewLevelRepositoryImpl,
	wire.Bind(new(repository.LevelRepository), new(*repository.LevelRepositoryImpl)),
	service.NewLevelServiceImpl,
	wire.Bind(new(service.LevelService), new(*service.LevelServiceImpl)),
	controller.NewLevelControllerImpl,
	wire.Bind(new(controller.LevelController), new(*controller.LevelControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		levelSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		NewServer,
	)
	return nil
}

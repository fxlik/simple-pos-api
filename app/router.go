package app

import (
	"fxlik/simple-post-api/controller"
	"fxlik/simple-post-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	categoryController controller.CategoryController,
	levelController controller.LevelController,
) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/levels", levelController.FindAll)
	router.GET("/api/levels/:levelId", levelController.FindById)
	router.POST("/api/levels", levelController.Create)
	router.PUT("/api/levels/:levelId", levelController.Update)
	router.DELETE("/api/levels/:levelId", levelController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}

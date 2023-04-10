package app

import (
	"fxlik/simple-post-api/controller"
	"fxlik/simple-post-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	categoryController controller.CategoryController,
	levelController controller.LevelController,
	supplierController controller.SupplierController,
	transactionStatusController controller.TransactionStatusController,
	transactionTypeController controller.TransactionTypeController,
	qtyTypeController controller.QtyTypeController,
	productController controller.ProductController,
) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/levels", levelController.FindAll)
	router.GET("/api/levels/:levelId", levelController.FindById)
	router.POST("/api/levels", levelController.Save)
	router.PUT("/api/levels/:levelId", levelController.Update)
	router.DELETE("/api/levels/:levelId", levelController.Delete)

	router.GET("/api/suppliers", supplierController.FindAll)
	router.GET("/api/suppliers/:supplierId", supplierController.FindById)
	router.POST("/api/suppliers", supplierController.Save)
	router.PUT("/api/suppliers/:supplierId", supplierController.Update)
	router.DELETE("/api/suppliers/:supplierId", supplierController.Delete)

	router.GET("/api/transaction-status", transactionStatusController.FindAll)
	router.GET("/api/transaction-status/:transactionStatusId", transactionStatusController.FindById)
	router.POST("/api/transaction-status", transactionStatusController.Save)
	router.PUT("/api/transaction-status/:transactionStatusId", transactionStatusController.Update)
	router.DELETE("/api/transaction-status/:transactionStatusId", transactionStatusController.Delete)

	router.GET("/api/transaction-type", transactionTypeController.FindAll)
	router.GET("/api/transaction-type/:transactionTypeId", transactionTypeController.FindById)
	router.POST("/api/transaction-type", transactionTypeController.Save)
	router.PUT("/api/transaction-type/:transactionTypeId", transactionTypeController.Update)
	router.DELETE("/api/transaction-type/:transactionTypeId", transactionTypeController.Delete)

	router.GET("/api/qty-type", qtyTypeController.FindAll)
	router.GET("/api/qty-type/:qtyTypeId", qtyTypeController.FindById)
	router.POST("/api/qty-type", qtyTypeController.Save)
	router.PUT("/api/qty-type/:qtyTypeId", qtyTypeController.Update)
	router.DELETE("/api/qty-type/:qtyTypeId", qtyTypeController.Delete)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Save)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}

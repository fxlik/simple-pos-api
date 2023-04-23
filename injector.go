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

var supplierSet = wire.NewSet(
	repository.NewSupplierRepositoryImpl,
	wire.Bind(new(repository.SupplierRepository), new(*repository.SupplierRepositoryImpl)),
	service.NewSupplierServiceImpl,
	wire.Bind(new(service.SupplierService), new(*service.SupplierServiceImpl)),
	controller.NewSupplierControllerImpl,
	wire.Bind(new(controller.SupplierController), new(*controller.SupplierControllerImpl)),
)

var transactionStatusSet = wire.NewSet(
	repository.NewTransactionStatusRepositoryImpl,
	wire.Bind(new(repository.TransactionStatusRepository), new(*repository.TransactionStatusRepositoryImpl)),
	service.NewTransactionStatusServiceImpl,
	wire.Bind(new(service.TransactionStatusService), new(*service.TransactionStatusServiceImpl)),
	controller.NewTransactionStatusControllerImpl,
	wire.Bind(new(controller.TransactionStatusController), new(*controller.TransactionStatusControllerImpl)),
)

var transactionTypeSet = wire.NewSet(
	repository.NewTransactionTypeRepositoryImpl,
	wire.Bind(new(repository.TransactionTypeRepository), new(*repository.TransactionTypeRepositoryImpl)),
	service.NewTransactionTypeServiceImpl,
	wire.Bind(new(service.TransactionTypeService), new(*service.TransactionTypeServiceImpl)),
	controller.NewTransactionTypeControllerImpl,
	wire.Bind(new(controller.TransactionTypeController), new(*controller.TransactionTypeControllerImpl)),
)

var qtyTypeSet = wire.NewSet(
	repository.NewQtyTypeRepositoryImpl,
	wire.Bind(new(repository.QtyTypeRepository), new(*repository.QtyTypeRepositoryImpl)),
	service.NewQtyTypeServiceImpl,
	wire.Bind(new(service.QtyTypeService), new(*service.QtyTypeServiceImpl)),
	controller.NewQtyTypeControllerImpl,
	wire.Bind(new(controller.QtyTypeController), new(*controller.QtyTypeControllerImpl)),
)

var productSet = wire.NewSet(
	repository.NewProductRepositoryImpl,
	wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)),
	service.NewProductServiceImpl,
	wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)),
	controller.NewProductControllerImpl,
	wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)),
)

var productPriceSet = wire.NewSet(
	repository.NewProductPriceRepositoryImpl,
	wire.Bind(new(repository.ProductPriceRepository), new(*repository.ProductPriceRepositoryImpl)),
	service.NewProductPriceServiceImpl,
	wire.Bind(new(service.ProductPriceService), new(*service.ProductPriceServiceImpl)),
	controller.NewProductPriceControllerImpl,
	wire.Bind(new(controller.ProductPriceController), new(*controller.ProductPriceControllerImpl)),
)

var transactionSet = wire.NewSet(
	repository.NewTransactionRepositoryImpl,
	wire.Bind(new(repository.TransactionRepository), new(*repository.TransactionRepositoryImpl)),
	service.NewTransactionServiceImpl,
	wire.Bind(new(service.TransactionService), new(*service.TransactionServiceImpl)),
	controller.NewTransactionControllerImpl,
	wire.Bind(new(controller.TransactionController), new(*controller.TransactionControllerImpl)),
)

var transactionItemSet = wire.NewSet(
	repository.NewTransactionItemRepositoryImpl,
	wire.Bind(new(repository.TransactionItemRepository), new(*repository.TransactionItemRepositoryImpl)),
	service.NewTransactionItemServiceImpl,
	wire.Bind(new(service.TransactionItemService), new(*service.TransactionItemServiceImpl)),
	controller.NewTransactionItemControllerImpl,
	wire.Bind(new(controller.TransactionItemController), new(*controller.TransactionItemControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		levelSet,
		supplierSet,
		transactionStatusSet,
		transactionTypeSet,
		qtyTypeSet,
		productSet,
		productPriceSet,
		transactionSet,
		transactionItemSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		NewServer,
	)
	return nil
}

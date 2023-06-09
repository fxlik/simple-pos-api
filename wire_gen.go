// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fxlik/simple-post-api/app"
	"fxlik/simple-post-api/controller"
	"fxlik/simple-post-api/repository"
	"fxlik/simple-post-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepositoryImpl()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceImpl := service.NewCategoryServiceImpl(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryControllerImpl(categoryServiceImpl)
	levelRepositoryImpl := repository.NewLevelRepositoryImpl()
	levelServiceImpl := service.NewLevelServiceImpl(levelRepositoryImpl, db, validate)
	levelControllerImpl := controller.NewLevelControllerImpl(levelServiceImpl)
	supplierRepositoryImpl := repository.NewSupplierRepositoryImpl()
	supplierServiceImpl := service.NewSupplierServiceImpl(supplierRepositoryImpl, db, validate)
	supplierControllerImpl := controller.NewSupplierControllerImpl(supplierServiceImpl)
	transactionStatusRepositoryImpl := repository.NewTransactionStatusRepositoryImpl()
	transactionStatusServiceImpl := service.NewTransactionStatusServiceImpl(transactionStatusRepositoryImpl, db, validate)
	transactionStatusControllerImpl := controller.NewTransactionStatusControllerImpl(transactionStatusServiceImpl)
	transactionTypeRepositoryImpl := repository.NewTransactionTypeRepositoryImpl()
	transactionTypeServiceImpl := service.NewTransactionTypeServiceImpl(transactionTypeRepositoryImpl, db, validate)
	transactionTypeControllerImpl := controller.NewTransactionTypeControllerImpl(transactionTypeServiceImpl)
	qtyTypeRepositoryImpl := repository.NewQtyTypeRepositoryImpl()
	qtyTypeServiceImpl := service.NewQtyTypeServiceImpl(qtyTypeRepositoryImpl, db, validate)
	qtyTypeControllerImpl := controller.NewQtyTypeControllerImpl(qtyTypeServiceImpl)
	productRepositoryImpl := repository.NewProductRepositoryImpl()
	productPriceRepositoryImpl := repository.NewProductPriceRepositoryImpl()
	productServiceImpl := service.NewProductServiceImpl(productRepositoryImpl, productPriceRepositoryImpl, db, validate)
	productControllerImpl := controller.NewProductControllerImpl(productServiceImpl)
	productPriceServiceImpl := service.NewProductPriceServiceImpl(productPriceRepositoryImpl, db, validate)
	productPriceControllerImpl := controller.NewProductPriceControllerImpl(productPriceServiceImpl)
	transactionRepositoryImpl := repository.NewTransactionRepositoryImpl()
	transactionServiceImpl := service.NewTransactionServiceImpl(transactionRepositoryImpl, db, validate)
	transactionControllerImpl := controller.NewTransactionControllerImpl(transactionServiceImpl)
	transactionItemRepositoryImpl := repository.NewTransactionItemRepositoryImpl()
	transactionItemServiceImpl := service.NewTransactionItemServiceImpl(transactionItemRepositoryImpl, transactionRepositoryImpl, productPriceRepositoryImpl, db, validate)
	transactionItemControllerImpl := controller.NewTransactionItemControllerImpl(transactionItemServiceImpl)
	router := app.NewRouter(categoryControllerImpl, levelControllerImpl, supplierControllerImpl, transactionStatusControllerImpl, transactionTypeControllerImpl, qtyTypeControllerImpl, productControllerImpl, productPriceControllerImpl, transactionControllerImpl, transactionItemControllerImpl)
	server := NewServer(router)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepositoryImpl, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)), service.NewCategoryServiceImpl, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryControllerImpl, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))

var levelSet = wire.NewSet(repository.NewLevelRepositoryImpl, wire.Bind(new(repository.LevelRepository), new(*repository.LevelRepositoryImpl)), service.NewLevelServiceImpl, wire.Bind(new(service.LevelService), new(*service.LevelServiceImpl)), controller.NewLevelControllerImpl, wire.Bind(new(controller.LevelController), new(*controller.LevelControllerImpl)))

var supplierSet = wire.NewSet(repository.NewSupplierRepositoryImpl, wire.Bind(new(repository.SupplierRepository), new(*repository.SupplierRepositoryImpl)), service.NewSupplierServiceImpl, wire.Bind(new(service.SupplierService), new(*service.SupplierServiceImpl)), controller.NewSupplierControllerImpl, wire.Bind(new(controller.SupplierController), new(*controller.SupplierControllerImpl)))

var transactionStatusSet = wire.NewSet(repository.NewTransactionStatusRepositoryImpl, wire.Bind(new(repository.TransactionStatusRepository), new(*repository.TransactionStatusRepositoryImpl)), service.NewTransactionStatusServiceImpl, wire.Bind(new(service.TransactionStatusService), new(*service.TransactionStatusServiceImpl)), controller.NewTransactionStatusControllerImpl, wire.Bind(new(controller.TransactionStatusController), new(*controller.TransactionStatusControllerImpl)))

var transactionTypeSet = wire.NewSet(repository.NewTransactionTypeRepositoryImpl, wire.Bind(new(repository.TransactionTypeRepository), new(*repository.TransactionTypeRepositoryImpl)), service.NewTransactionTypeServiceImpl, wire.Bind(new(service.TransactionTypeService), new(*service.TransactionTypeServiceImpl)), controller.NewTransactionTypeControllerImpl, wire.Bind(new(controller.TransactionTypeController), new(*controller.TransactionTypeControllerImpl)))

var qtyTypeSet = wire.NewSet(repository.NewQtyTypeRepositoryImpl, wire.Bind(new(repository.QtyTypeRepository), new(*repository.QtyTypeRepositoryImpl)), service.NewQtyTypeServiceImpl, wire.Bind(new(service.QtyTypeService), new(*service.QtyTypeServiceImpl)), controller.NewQtyTypeControllerImpl, wire.Bind(new(controller.QtyTypeController), new(*controller.QtyTypeControllerImpl)))

var productSet = wire.NewSet(repository.NewProductRepositoryImpl, wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)), service.NewProductServiceImpl, wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)), controller.NewProductControllerImpl, wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)))

var productPriceSet = wire.NewSet(repository.NewProductPriceRepositoryImpl, wire.Bind(new(repository.ProductPriceRepository), new(*repository.ProductPriceRepositoryImpl)), service.NewProductPriceServiceImpl, wire.Bind(new(service.ProductPriceService), new(*service.ProductPriceServiceImpl)), controller.NewProductPriceControllerImpl, wire.Bind(new(controller.ProductPriceController), new(*controller.ProductPriceControllerImpl)))

var transactionSet = wire.NewSet(repository.NewTransactionRepositoryImpl, wire.Bind(new(repository.TransactionRepository), new(*repository.TransactionRepositoryImpl)), service.NewTransactionServiceImpl, wire.Bind(new(service.TransactionService), new(*service.TransactionServiceImpl)), controller.NewTransactionControllerImpl, wire.Bind(new(controller.TransactionController), new(*controller.TransactionControllerImpl)))

var transactionItemSet = wire.NewSet(repository.NewTransactionItemRepositoryImpl, wire.Bind(new(repository.TransactionItemRepository), new(*repository.TransactionItemRepositoryImpl)), service.NewTransactionItemServiceImpl, wire.Bind(new(service.TransactionItemService), new(*service.TransactionItemServiceImpl)), controller.NewTransactionItemControllerImpl, wire.Bind(new(controller.TransactionItemController), new(*controller.TransactionItemControllerImpl)))

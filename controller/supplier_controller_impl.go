package controller

import (
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type SupplierControllerImpl struct {
	SupplierService service.SupplierService
}

func NewSupplierControllerImpl(supplierService service.SupplierService) *SupplierControllerImpl {
	return &SupplierControllerImpl{
		SupplierService: supplierService,
	}
}

func (controller *SupplierControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	supplierCreateRequest := web.SupplierCreateRequest{}
	helper.ReadFromRequestBody(request, &supplierCreateRequest)

	supplierResponse := controller.SupplierService.Save(request.Context(), supplierCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   supplierResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	supplierUpdateRequest := web.SupplierUpdateRequest{}
	helper.ReadFromRequestBody(request, &supplierUpdateRequest)
	supplierId := params.ByName("supplierId")
	id, err := strconv.Atoi(supplierId)
	helper.PanicIfError(err)

	supplierUpdateRequest.Id = int32(id)
	supplierResponse := controller.SupplierService.Update(request.Context(), supplierUpdateRequest)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   supplierResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	supplierId := params.ByName("supplierId")
	id, err := strconv.Atoi(supplierId)
	helper.PanicIfError(err)
	controller.SupplierService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	supplierId := params.ByName("supplierId")
	id, err := strconv.Atoi(supplierId)
	helper.PanicIfError(err)

	supplierResponse := controller.SupplierService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   supplierResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	supplierResponses := controller.SupplierService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   supplierResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

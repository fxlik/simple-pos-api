package controller

import (
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type TransactionTypeControllerImpl struct {
	TransactionTypeService service.TransactionTypeService
}

func NewTransactionTypeControllerImpl(transactionTypeService service.TransactionTypeService) *TransactionTypeControllerImpl {
	return &TransactionTypeControllerImpl{
		TransactionTypeService: transactionTypeService,
	}
}

func (controller *TransactionTypeControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionTypeCreateRequest := web.TransactionTypeCreateRequest{}
	helper.ReadFromRequestBody(request, &transactionTypeCreateRequest)

	transactionTypeResponse := controller.TransactionTypeService.Save(request.Context(), transactionTypeCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionTypeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionTypeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionTypeUpdateRequest := web.TransactionTypeUpdateRequest{}
	helper.ReadFromRequestBody(request, &transactionTypeUpdateRequest)

	transactionTypeId := params.ByName("transactionTypeId")
	id, err := strconv.Atoi(transactionTypeId)
	helper.PanicIfError(err)

	transactionTypeUpdateRequest.Id = int32(id)
	transactionTypeResponse := controller.TransactionTypeService.Update(request.Context(), transactionTypeUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionTypeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionTypeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionTypeId := params.ByName("transactionTypeId")
	id, err := strconv.Atoi(transactionTypeId)
	helper.PanicIfError(err)

	controller.TransactionTypeService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionTypeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionTypeId := params.ByName("transactionTypeId")
	id, err := strconv.Atoi(transactionTypeId)
	helper.PanicIfError(err)

	transactionTypeResponse := controller.TransactionTypeService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionTypeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionTypeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionTypeResponses := controller.TransactionTypeService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionTypeResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

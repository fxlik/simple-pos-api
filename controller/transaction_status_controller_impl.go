package controller

import (
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type TransactionStatusControllerImpl struct {
	TransactionStatusService service.TransactionStatusService
}

func NewTransactionStatusControllerImpl(transactionStatusService service.TransactionStatusService) *TransactionStatusControllerImpl {
	return &TransactionStatusControllerImpl{
		TransactionStatusService: transactionStatusService,
	}
}

func (controller *TransactionStatusControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionStatusCreateRequest := web.TransactionStatusCreateRequest{}
	helper.ReadFromRequestBody(request, &transactionStatusCreateRequest)

	transactionStatusResponse := controller.TransactionStatusService.Save(request.Context(), transactionStatusCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionStatusResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionStatusControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionStatusUpdateRequest := web.TransactionStatusUpdateRequest{}
	helper.ReadFromRequestBody(request, &transactionStatusUpdateRequest)
	transactionStatusId := params.ByName("transactionStatusId")
	id, err := strconv.Atoi(transactionStatusId)
	helper.PanicIfError(err)

	transactionStatusUpdateRequest.Id = int32(id)
	transactionStatusResponse := controller.TransactionStatusService.Update(request.Context(), transactionStatusUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionStatusResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionStatusControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionStatusId := params.ByName("transactionStatusId")
	id, err := strconv.Atoi(transactionStatusId)
	helper.PanicIfError(err)
	controller.TransactionStatusService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionStatusControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionStatusId := params.ByName("transactionStatusId")
	id, err := strconv.Atoi(transactionStatusId)
	helper.PanicIfError(err)
	transactionStatusResponse := controller.TransactionStatusService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionStatusResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionStatusControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionStatusResponses := controller.TransactionStatusService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionStatusResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

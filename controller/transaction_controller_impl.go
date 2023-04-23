package controller

import (
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionControllerImpl(transactionService service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionCreateRequest := web.TransactionCreateRequest{}

	contentType := request.Header.Get("Content-Type")
	helper.CheckContentType(contentType, request, &transactionCreateRequest)

	transactionResponse := controller.TransactionService.Save(request.Context(), transactionCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionUpdateRequest := web.TransactionUpdateRequest{}
	contentType := request.Header.Get("Content-Type")
	helper.CheckContentType(contentType, request, &transactionUpdateRequest)

	transactionId := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionId)
	helper.PanicIfError(err)

	transactionUpdateRequest.Id = int32(id)
	transactionResponse := controller.TransactionService.Update(request.Context(), transactionUpdateRequest)

	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionId := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionId)
	helper.PanicIfError(err)
	controller.TransactionService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionId := params.ByName("transactionId")
	id, err := strconv.Atoi(transactionId)
	helper.PanicIfError(err)

	transactionResponse := controller.TransactionService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionResponses := controller.TransactionService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

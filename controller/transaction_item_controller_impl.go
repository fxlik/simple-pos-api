package controller

import (
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type TransactionItemControllerImpl struct {
	TransactionItemService service.TransactionItemService
}

func NewTransactionItemControllerImpl(transactionItemService service.TransactionItemService) *TransactionItemControllerImpl {
	return &TransactionItemControllerImpl{
		TransactionItemService: transactionItemService,
	}
}

func (controller *TransactionItemControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionItemCreateRequest := web.TransactionItemCreateRequest{}
	contentType := request.Header.Get("Content-Type")
	helper.CheckContentType(contentType, request, &transactionItemCreateRequest)

	transactionItemResponse := controller.TransactionItemService.Save(request.Context(), transactionItemCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionItemResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionItemControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionItemUpdateRequest := web.TransactionItemUpdateRequest{}
	contentType := request.Header.Get("Content-Type")
	helper.CheckContentType(contentType, request, &transactionItemUpdateRequest)

	transactionItemId := params.ByName("transactionItemId")
	id, errConv := strconv.Atoi(transactionItemId)
	helper.PanicIfError(errConv)
	transactionItemUpdateRequest.Id = int32(id)
	transactionResponse := controller.TransactionItemService.Update(request.Context(), transactionItemUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionItemControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionItemId := params.ByName("transactionItemId")
	id, errConv := strconv.Atoi(transactionItemId)
	helper.PanicIfError(errConv)
	controller.TransactionItemService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionItemControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionItemId := params.ByName("transactionItemId")
	id, errConv := strconv.Atoi(transactionItemId)
	helper.PanicIfError(errConv)

	transactionItemResponse := controller.TransactionItemService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionItemResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionItemControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionItemResponses := controller.TransactionItemService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionItemResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransactionItemControllerImpl) FindAllByTransactionId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionId := params.ByName("transactionId")
	id, errConv := strconv.Atoi(transactionId)
	helper.PanicIfError(errConv)

	transactionItemResponses := controller.TransactionItemService.FindAllByTransactionId(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   transactionItemResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

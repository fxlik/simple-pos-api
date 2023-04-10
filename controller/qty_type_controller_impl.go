package controller

import (
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

type QtyTypeControllerImpl struct {
	QtyTypeService service.QtyTypeService
}

func NewQtyTypeControllerImpl(qtyTypeService service.QtyTypeService) *QtyTypeControllerImpl {
	return &QtyTypeControllerImpl{
		QtyTypeService: qtyTypeService,
	}
}

func (controller *QtyTypeControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	qtyTypeCreateRequest := web.QtyTypeCreateRequest{}
	contentType := request.Header.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "multipart/form-data"):
		helper.ReadFromRequestMultipartFormData(request, &qtyTypeCreateRequest)
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		helper.ReadFromRequestForm(request, &qtyTypeCreateRequest)
	default:
		//if content type is "application/json"
		helper.ReadFromRequestBody(request, &qtyTypeCreateRequest)
	}
	qtyTypeResponse := controller.QtyTypeService.Save(request.Context(), qtyTypeCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   qtyTypeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QtyTypeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	qtyTypeUpdateRequest := web.QtyTypeUpdateRequest{}
	helper.ReadFromRequestBody(request, &qtyTypeUpdateRequest)

	qtyTypeId := params.ByName("qtyTypeId")
	id, err := strconv.Atoi(qtyTypeId)
	helper.PanicIfError(err)

	qtyTypeUpdateRequest.Id = int32(id)
	qtyTypeResponse := controller.QtyTypeService.Update(request.Context(), qtyTypeUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   qtyTypeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QtyTypeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	qtyTypeId := params.ByName("qtyTypeId")
	id, err := strconv.Atoi(qtyTypeId)
	helper.PanicIfError(err)
	controller.QtyTypeService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QtyTypeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	qtyTypeId := params.ByName("qtyTypeId")
	id, err := strconv.Atoi(qtyTypeId)
	helper.PanicIfError(err)
	qtyTypeResponse := controller.QtyTypeService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   qtyTypeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QtyTypeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	qtyTypeResponses := controller.QtyTypeService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   qtyTypeResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

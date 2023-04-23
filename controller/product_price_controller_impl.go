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

type ProductPriceControllerImpl struct {
	ProductPriceService service.ProductPriceService
}

func NewProductPriceControllerImpl(productPriceService service.ProductPriceService) *ProductPriceControllerImpl {
	return &ProductPriceControllerImpl{
		ProductPriceService: productPriceService,
	}
}

func (controller *ProductPriceControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productPriceCreateRequest := web.ProductPriceCreateRequest{}

	contentType := request.Header.Get("Content-Type")
	helper.CheckContentType(contentType, request, &productPriceCreateRequest)

	productPriceResponse := controller.ProductPriceService.Save(request.Context(), productPriceCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productPriceResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductPriceControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productPriceUpdateRequest := web.ProductPriceUpdateRequest{}
	contentType := request.Header.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "multipart/form-data"):
		helper.ReadFromRequestMultipartFormData(request, &productPriceUpdateRequest)
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		helper.ReadFromRequestForm(request, &productPriceUpdateRequest)
	default:
		//if content type is "application/json"
		helper.ReadFromRequestBody(request, &productPriceUpdateRequest)
	}

	productPriceId := params.ByName("productPriceId")
	id, err := strconv.Atoi(productPriceId)
	helper.PanicIfError(err)
	productPriceUpdateRequest.ProductId = int32(id)

	productPriceResponse := controller.ProductPriceService.Update(request.Context(), productPriceUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productPriceResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductPriceControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productPriceId := params.ByName("productPriceId")
	id, err := strconv.Atoi(productPriceId)
	helper.PanicIfError(err)
	controller.ProductPriceService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductPriceControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productPriceId := params.ByName("productPriceId")
	id, err := strconv.Atoi(productPriceId)
	helper.PanicIfError(err)
	productPriceResponse := controller.ProductPriceService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productPriceResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductPriceControllerImpl) FindOneByProductId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	productPriceResponse := controller.ProductPriceService.FindOneByProductId(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productPriceResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductPriceControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productPriceResponses := controller.ProductPriceService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productPriceResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

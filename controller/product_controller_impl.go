package controller

import (
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/h2non/filetype"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductControllerImpl(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}

	//file upload handing
	var fileName string
	file, fileHeader, err := request.FormFile("image")
	helper.PanicIfError(err)
	defer file.Close()
	fileByte, err := io.ReadAll(file)
	helper.PanicIfError(err)

	if filetype.IsImage(fileByte) {
		fileName = "resources/uploads/" + "product-" + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(fileHeader.Filename)
		err := os.WriteFile(fileName, fileByte, 0666)
		helper.PanicIfError(err)

		//cara yang dikomen dibawah ini, sebaiknya menggunakan nama file asli, karena hanya mengkopi data
		//dari satu direktori ke direktori lain (direktori project)

		//fileDestination, err := os.Create(fileName)
		//helper.PanicIfError(err)
		//
		//_, errCopy := io.Copy(fileDestination, file)
		//helper.PanicIfError(errCopy)
	} else {
		newError := errors.New("file type is not supported")
		helper.PanicIfError(newError)
	}
	productCreateRequest.Image = fileName

	contentType := request.Header.Get("Content-Type")
	helper.CheckContentType(contentType, request, &productCreateRequest)

	productResponse := controller.ProductService.Save(request.Context(), productCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}

	//file upload handing
	var fileName string
	file, handler, err := request.FormFile("image")
	helper.PanicIfError(err)
	defer file.Close()

	fileByte, err := io.ReadAll(file)
	helper.PanicIfError(err)

	if filetype.IsImage(fileByte) {
		fileName = "resources/uploads/" + "product-" + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(handler.Filename)
		err := os.WriteFile(fileName, fileByte, 0666)
		helper.PanicIfError(err)
	} else {
		newError := errors.New("file type is not supported")
		helper.PanicIfError(newError)
	}

	contentType := request.Header.Get("Content-Type")
	helper.CheckContentType(contentType, request, &productUpdateRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	productUpdateRequest.Id = int32(id)
	productUpdateRequest.Image = fileName

	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	controller.ProductService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	productResponse := controller.ProductService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

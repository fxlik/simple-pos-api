package controller

import (
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/web"
	"fxlik/simple-post-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type LevelControllerImpl struct {
	LevelService service.LevelService
}

func NewLevelControllerImpl(levelService service.LevelService) *LevelControllerImpl {
	return &LevelControllerImpl{
		LevelService: levelService,
	}
}

func (controller *LevelControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	levelCreateRequest := web.LevelCreateRequest{}
	helper.ReadFromRequestBody(request, &levelCreateRequest)

	levelResponse := controller.LevelService.Create(request.Context(), levelCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   levelResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LevelControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	levelUpdateRequest := web.LevelUpdateRequest{}
	helper.ReadFromRequestBody(request, &levelUpdateRequest)
	levelId := params.ByName("levelId")
	id, err := strconv.Atoi(levelId)
	helper.PanicIfError(err)

	levelUpdateRequest.Id = int32(id)
	levelResponse := controller.LevelService.Update(request.Context(), levelUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   levelResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LevelControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	levelId := params.ByName("levelId")
	id, err := strconv.Atoi(levelId)
	helper.PanicIfError(err)
	controller.LevelService.Delete(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LevelControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	levelId := params.ByName("levelId")
	id, err := strconv.Atoi(levelId)
	helper.PanicIfError(err)
	levelResponse := controller.LevelService.FindById(request.Context(), int32(id))
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   levelResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LevelControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	levelResponses := controller.LevelService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   levelResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

package helper

import (
	"encoding/json"
	"github.com/go-playground/form"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

var decoder *form.Decoder

func ReadFromRequestForm(request *http.Request, result interface{}) {
	err := request.ParseForm()
	PanicIfError(err)
	decoder = form.NewDecoder()
	values := request.PostForm
	errDecode := decoder.Decode(&result, values)
	PanicIfError(errDecode)
}

func ReadFromRequestMultipartFormData(request *http.Request, result interface{}) {
	err := request.ParseMultipartForm(32 << 20)
	PanicIfError(err)
	request.Header.Set("Content-Type", "multipart/form-data")
	decoder = form.NewDecoder()
	values := request.PostForm
	errDecode := decoder.Decode(&result, values)
	PanicIfError(errDecode)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

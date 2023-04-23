package helper

import (
	"net/http"
	"strings"
)

func CheckContentType(contentType string, request *http.Request, entityStruct interface{}) {
	contentType = request.Header.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "multipart/form-data"):
		ReadFromRequestMultipartFormData(request, &entityStruct)
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		ReadFromRequestForm(request, &entityStruct)
	default:
		//if content type is "application/json"
		ReadFromRequestBody(request, &entityStruct)
	}
}

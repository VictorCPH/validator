package validator

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MultipartMemory is the maximum permitted size of the request body in an HTTP request.
// The whole request body is parsed and up to a total of MultipartMemory bytes of its file
// parts are stored in memory, with the remainder stored on disk in temporary files.
var MultipartMemory int64 = 64 * 1024 * 1024

// Bind takes data out of the request and deserializes into a interface obj according
// to the Content-Type of the request. If no Content-Type is specified, there
// better be data in the query string, otherwise an error will be produced.
//
// A non-nil return value may be an Errors value.
func Bind(req *http.Request, obj interface{}) error {
	method := req.Method
	contentType := filterFlags(req.Header.Get("Content-Type"))

	if method == "GET" {
		return BindURL(req, obj)
	}

	switch contentType {
	case ContentTypeJson:
		return BindJson(req, obj)
	case ContentTypeMultipart:
		return BindMultipart(req, obj)
	case ContentTypeForm:
		return BindForm(req, obj)
	default:
		return fmt.Errorf(ERR_UNSUPPORTED_CONTENT_TYPE)
	}
}

func BindForm(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return fmt.Errorf("%v: %v", ERR_PARSE_FORM, err.Error())
	}
	if err := coerce(obj, req.Form, nil); err != nil {
		return err
	}
	return validate(obj)
}

func BindMultipart(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(MultipartMemory); err != nil {
		return fmt.Errorf("%v: %v", ERR_PARSE_MULTIPART_FORM, err.Error())
	}
	if err := coerce(obj, req.Form, req.MultipartForm.File); err != nil {
		return err
	}
	return validate(obj)
}

func BindURL(req *http.Request, obj interface{}) error {
	if err := coerce(obj, req.URL.Query(), nil); err != nil {
		return err
	}
	return validate(obj)
}

func BindJson(req *http.Request, obj interface{}) error {
	err := json.NewDecoder(req.Body).Decode(obj)
	if err != nil {
		return fmt.Errorf("%v: %v", ERR_DECODE_JSON, err.Error())
	}
	return validate(obj)
}

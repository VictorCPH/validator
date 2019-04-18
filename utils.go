package validator

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

// Check str is in values
func isIn(str string, values []string) bool {
	for _, value := range values {
		if str == value {
			return true
		}
	}

	return false
}

func request(method, path, body, contentType string) *http.Request {
	req, err := http.NewRequest(method, path, bytes.NewBufferString(body))
	if err != nil {
		panic(err)
	}
	if contentType != "" {
		req.Header.Add("Content-type", contentType)
	}
	return req
}

func requestMultipartForm(path string, params map[string]string, files map[string]string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for name, path := range files {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		part, err := writer.CreateFormFile(name, path)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(part, f)
		if err != nil {
			panic(err)
		}
	}

	for k, v := range params {
		err := writer.WriteField(k, v)
		if err != nil {
			panic(err)
		}
	}
	err := writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

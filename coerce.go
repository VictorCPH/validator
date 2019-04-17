package validator

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"strconv"
)

// coerce tries to set the value with the type of the param. If fail then return error
func coerce(obj interface{}, formData map[string][]string, formFile map[string][]*multipart.FileHeader) error {
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag
		name := tag.Get("form")
		field := val.Field(i)

		err := coerceField(field, name, tag, formData, formFile)
		if err != nil {
			if err.Error() == ERR_OPTIONAL_PARAM_NOT_FOUND {
				continue
			} else {
				return fmt.Errorf("%s: %s", name, err.Error())
			}
		}
	}
	return nil
}

func coerceField(val reflect.Value, name string, tag reflect.StructTag, formData map[string][]string,
	formFile map[string][]*multipart.FileHeader) (err error) {
	var params []string
	var files []*multipart.FileHeader
	notFound := false

	// Get params and files
	if vs := formData[name]; len(vs) > 0 {
		params = vs
	} else if vs := formData[name+"[]"]; len(vs) > 0 {
		params = vs
	} else if formFile != nil && len(formFile[name]) > 0 {
		files = formFile[name]
	} else {
		notFound = true
	}

	// Check exist
	if notFound {
		switch tag.Get("valid") {
		case "required":
			return fmt.Errorf(ERR_PARAM_NOT_FOUND)
		case "optional":
			if len(tag.Get("default")) != 0 {
				params = []string{tag.Get("default")}
			} else {
				return fmt.Errorf(ERR_OPTIONAL_PARAM_NOT_FOUND)
			}
		}
	}

	if tag.Get("type") == "file" {
		// Read file to bytes
		if len(files) > 0 {
			f, err := files[0].Open()
			if err != nil {
				return fmt.Errorf(ERR_CORRUPTED_FILE)
			}
			blob, err := ioutil.ReadAll(f)
			if err != nil {
				return fmt.Errorf(ERR_CORRUPTED_FILE)
			}
			val.SetBytes(blob)
		} else if len(params) > 0 {
			return fmt.Errorf(ERR_FILE_TYPE_INVALID)
		}
	} else if tag.Get("type") == "base64" {
		// Decode base64 string to bytes
		decoded, err := base64.StdEncoding.DecodeString(params[0])
		if err != nil {
			return fmt.Errorf(ERR_INVALID_BASE64)
		}
		val.SetBytes(decoded)
	} else {
		// Other type
		switch val.Kind() {
		case reflect.Slice:
			s := reflect.MakeSlice(val.Type(), len(params), len(params))
			for i, v := range params {
				err = setValue(s.Index(i), v)
				if err != nil {
					return fmt.Errorf(ERR_PARAM_INVALID, val.Kind().String())
				}
			}
			val.Set(s)
		default:
			setValue(val, params[0])
			if err != nil {
				return fmt.Errorf(ERR_PARAM_INVALID, val.Kind().String())
			}
		}
	}
	return nil
}

func setValue(val reflect.Value, param string) error {
	switch val.Kind() {
	case reflect.Int:
		i, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return err
		}
		val.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(param)
		if err != nil {
			return err
		}
		val.SetBool(b)
	case reflect.Float32:
		f32, err := strconv.ParseFloat(param, 32)
		if err != nil {
			return err
		}
		val.SetFloat(f32)
	case reflect.Float64:
		f64, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return err
		}
		val.SetFloat(f64)
	case reflect.String:
		val.SetString(param)
	default:
	}
	return nil
}

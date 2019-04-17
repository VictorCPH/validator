package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Validate check the value of the param by tag. If not valid then return error
func validate(obj interface{}) error {
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag
		name := tag.Get("form")
		field := val.Field(i)

		if err := validateField(field, tag); err != nil {
			return fmt.Errorf("%s: %s", name, err.Error())
		}
	}
	return nil
}

// TODO: more check
func validateField(v reflect.Value, tag reflect.StructTag) (err error) {
	switch v.Kind() {
	case reflect.String:
		if !utf8.Valid([]byte(v.String())) {
			return fmt.Errorf(ERR_INVALID_UTF8_STRING)
		}
		if len(v.String()) == 0 {
			return fmt.Errorf(ERR_BLANK_STRING)
		}
		if len(tag.Get("values")) != 0 {
			values := strings.Split(tag.Get("values"), "|")
			if !isIn(v.String(), values) {
				return fmt.Errorf(ERR_INVALID_ENUMERATION, v.String(), values)
			}
		}
		if len(tag.Get("regexp")) != 0 {
			re := regexp.MustCompile(tag.Get("regexp"))
			if !re.Match([]byte(v.String())) {
				return fmt.Errorf(ERR_WRONG_FORMAT, re)
			}
		}
	case reflect.Slice:
		var typeOfBytes = reflect.TypeOf([]byte(nil))
		if tag.Get("type") == "file" && v.Type() == typeOfBytes {
			if len(tag.Get("max_size")) != 0 {
				max_size, err := strconv.Atoi(tag.Get("max_size"))
				if err != nil {
					fmt.Errorf(ERR_INVALID_MAX_SIZE)
				}
				if len(v.Bytes()) > max_size {
					return fmt.Errorf(ERR_PARAM_FILE_TOO_LARGE, max_size)
				}
			}
		}
	case reflect.Int:
	}
	return nil
}

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
func validate(obj interface{}, format string) error {
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag
		name := tag.Get(format)
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
			if !re.MatchString(v.String()) {
				return fmt.Errorf(ERR_WRONG_FORMAT, re)
			}
		}
	case reflect.Slice:
		// []byte
		var typeOfBytes = reflect.TypeOf([]byte(nil))
		if tag.Get("type") == "file" && v.Type() == typeOfBytes {
			if len(tag.Get("max_size")) != 0 {
				max_size, err := strconv.ParseInt(tag.Get("max_size"), 10, 64)
				if err != nil {
					return fmt.Errorf(ERR_INVALID_MAX_SIZE_TAG)
				}
				if len(v.Bytes()) > int(max_size) {
					return fmt.Errorf(ERR_PARAM_FILE_TOO_LARGE, max_size)
				}
			}
		}
		// Other
		for i := 0; i < v.Len(); i++ {
			err = validateField(v.Index(i), tag)
			if err != nil {
				return err
			}
		}
	case reflect.Int:
		if len(tag.Get("max")) != 0 {
			max, err := strconv.ParseInt(tag.Get("max"), 10, 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_MAX_TAG)
			}
			if v.Int() > max {
				return fmt.Errorf(ERR_GREATER_THAN_MAX, tag.Get("max"))
			}
		}
		if len(tag.Get("min")) != 0 {
			min, err := strconv.ParseInt(tag.Get("min"), 10, 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_MIN_TAG)
			}
			if v.Int() < min {
				return fmt.Errorf(ERR_SMALLER_THAN_MIN, tag.Get("min"))
			}
		}
		if len(tag.Get("range")) != 0 {
			r := strings.Split(tag.Get("range"), "|")
			min, err := strconv.ParseInt(r[0], 10, 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_RANGE_TAG)
			}
			max, err := strconv.ParseInt(r[1], 10, 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_RANGE_TAG)
			}
			if v.Int() < min || v.Int() > max {
				return fmt.Errorf(ERR_NOT_IN_RANGE, r[0], r[1])
			}
		}
	case reflect.Float32, reflect.Float64:
		if len(tag.Get("max")) != 0 {
			max, err := strconv.ParseFloat(tag.Get("max"), 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_MAX_TAG)
			}
			if v.Float() > max {
				return fmt.Errorf(ERR_GREATER_THAN_MAX, tag.Get("max"))
			}
		}
		if len(tag.Get("min")) != 0 {
			min, err := strconv.ParseFloat(tag.Get("min"), 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_MIN_TAG)
			}
			if v.Float() < min {
				return fmt.Errorf(ERR_SMALLER_THAN_MIN, tag.Get("min"))
			}
		}
		if len(tag.Get("range")) != 0 {
			r := strings.Split(tag.Get("range"), "|")
			min, err := strconv.ParseFloat(r[0], 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_RANGE_TAG)
			}
			max, err := strconv.ParseFloat(r[1], 64)
			if err != nil {
				return fmt.Errorf(ERR_INVALID_RANGE_TAG)
			}
			if v.Float() < min || v.Float() > max {
				return fmt.Errorf(ERR_NOT_IN_RANGE, r[0], r[1])
			}
		}
	}
	return nil
}

package validator

// Content Type
const (
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeMultipart = "multipart/form-data"
	ContentTypeJson      = "application/json"
)

// Error Message
const (
	// Parse data error
	ERR_EMPTY_CONTENT_TYPE       = "empty Content-Type"
	ERR_UNSUPPORTED_CONTENT_TYPE = "unsupported Content-Type"
	ERR_PARSE_FORM               = "parse form failed"
	ERR_PARSE_MULTIPART_FORM     = "parse multipart form failed"
	ERR_DECODE_JSON              = "decode json failed"

	// Coerce error
	ERR_OPTIONAL_PARAM_NOT_FOUND = "optional param not found"
	ERR_PARAM_NOT_FOUND          = "not found"
	ERR_PARAM_INVALID            = "%s expected"
	ERR_CORRUPTED_FILE           = "corrupted file"
	ERR_PARAM_FILE_NOT_FOUND     = "file not found"
	ERR_FILE_TYPE_INVALID        = "file expected"

	// Validate error
	ERR_PARAM_FILE_TOO_LARGE = "file larger than %d bytes"
	ERR_INVALID_MAX_SIZE_TAG = "invalid `max_size` tag, must be int"
	ERR_INVALID_VALID_TAG    = "invalid `valid` tag, must be `required` or `optional`"
	ERR_INVALID_MAX_TAG      = "invalid `max` tag, must be int or float"
	ERR_INVALID_MIN_TAG      = "invalid `min` tag, must be int or float"
	ERR_INVALID_RANGE_TAG    = "invalid `range` tag, must be (int|int) or (float|float)"
	ERR_INVALID_BASE64       = "invalid base64 string"
	ERR_INVALID_UTF8_STRING  = "invalid utf8 string"
	ERR_GREATER_THAN_MAX     = "greater than %s"
	ERR_SMALLER_THAN_MIN     = "smaller than %s"
	ERR_BLANK_STRING         = "blank string"
	ERR_INVALID_ENUMERATION  = "%s is not in %s"
	ERR_WRONG_FORMAT         = "wrong format, shold match regexp `%s`"
	ERR_NOT_IN_RANGE         = "not in range (%s, %s)"
)

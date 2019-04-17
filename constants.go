package validator

// Content Type
const (
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeMultipart = "multipart/form-data"
	ContentTypeJson      = "application/json"
)

// Error Message
const (
	ERR_EMPTY_CONTENT_TYPE       = "empty Content-Type"
	ERR_UNSUPPORTED_CONTENT_TYPE = "unsupported Content-Type"
	ERR_PARSE_FORM               = "parse form failed"
	ERR_PARSE_MULTIPART_FORM     = "parse multipart form failed"
	ERR_DECODE_JSON              = "decode json failed"
	ERR_OPTIONAL_PARAM_NOT_FOUND = "optional param not found"
	ERR_PARAM_NOT_FOUND          = "not found"
	ERR_PARAM_INVALID            = "%s expected"
	ERR_CORRUPTED_FILE           = "corrupted file"
	ERR_PARAM_FILE_NOT_FOUND     = "file not found"
	ERR_FILE_TYPE_INVALID        = "file expected"
	ERR_INVALID_MAX_SIZE         = "invalid max_size, must be int"
	ERR_PARAM_FILE_TOO_LARGE     = "file larger than %d bytes"
	ERR_INVALID_BASE64           = "invalid base64 string"
	ERR_INVALID_UTF8_STRING      = "invalid utf8 string"
	ERR_BLANK_STRING             = "blank string"
	ERR_INVALID_ENUMERATION      = "%s is not in %s"
	ERR_INVALID_RANGE            = "invalid range"
	ERR_WRONG_FORMAT             = "wrong format, shold match regexp `%s`"
	ERR_NOT_IN_RANGE             = "not in range (%s, %s)"
)

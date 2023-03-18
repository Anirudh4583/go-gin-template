package e

var ErrorMessageFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "request parameter error",

	NOT_FOUND: 						 "Not Found",

	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token has timed out",
	ERROR_AUTH_TOKEN:                "Token generation failed",
	ERROR_AUTH:                      "Token error",
}

func GetErrorMessage(code int) string {
	msg, ok := ErrorMessageFlags[code]
	if ok {
		return msg
	}

	return ErrorMessageFlags[ERROR]
}
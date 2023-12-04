package internal

import (
	"net/http"
	pkg_constant "osstp-go-hive/pkg/constant"
)

func ResponseCode(httpCode int, msgCode int, message string) (int, string) {
	var (
		code int
		msg  string
	)

	switch httpCode {
	case http.StatusOK:
		// Only success return 1
		code = pkg_constant.Success
	default:
		// Others return 0
		code = pkg_constant.Fail
	}

	// Default msgCode
	// if temp == "", use message value
	if temp := pkg_constant.StatusMsg(msgCode); temp != "" {
		msg = temp
	} else {
		msg = message
	}

	return code, msg
}

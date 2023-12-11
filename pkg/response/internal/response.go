package internal

import (
	pkg_constant "osstp-go-hive/pkg/constant"
)

func ResponseCode(httpCode int, msgCode int) (int, int) {
	var (
		code        int
		tempMsgCode int
	)

	switch httpCode {
	case 200:
		// Only success return 1
		code = pkg_constant.Success
		tempMsgCode = 1
	default:
		// Others return 0
		code = pkg_constant.Fail
		tempMsgCode = msgCode
	}

	return code, tempMsgCode
}

func GetOk(msg string) string {
	var (
		ok string = ""
	)

	if msg == "" {
		ok = "OK"
	} else {
		ok = msg
	}

	return ok
}

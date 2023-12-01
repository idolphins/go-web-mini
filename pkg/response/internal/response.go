package internal

import (
	"net/http"
)

func ResponseCode(httpCode int) int {
	var (
		code int
	)

	switch httpCode {
	case http.StatusOK:
		code = 1
	default:
		code = 0
	}

	// switch pkgCode.(type) {
	// case int:
	// 	msg = pkg_constant.StatusMsg(int(pkgCode.(int)))
	// 	if msg == "" {
	// 		msg = message
	// 	}
	// case string:
	// 	msg = fmt.Sprintf("%v", pkgCode)
	// }

	return code
}

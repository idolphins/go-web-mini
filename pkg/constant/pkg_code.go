package pkg_constant

const (
	Success             = 1
	Fail                = 0
	Invalid_params      = 3
	Error_database_CRUD = 3306

	Error_user_exist     = 10001
	Error_user_not_exist = 10002
	Error_password       = 10003
	Error_captcha        = 10004
	Error_logout         = 10005

	// //token相关
	// ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	// ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	// ERROR_AUTH_TOKEN               = 20003
	// ERROR_AUTH                     = 20004
	// ERROR_AUTH_CHECK_FAIL          = 20005

	// ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	// ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	// ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

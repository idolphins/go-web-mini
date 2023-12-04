package pkg_constant

const (
	Success              = 1
	Fail                 = 0
	Error_invalid_params = 3

	Error_database_CRUD = 33060

	Error_user_exist     = 10001
	Error_user_not_exist = 10002
	Error_user_password  = 10003
	Error_user_captcha   = 10004
	Error_user_logout    = 10005

	Error_token_check   = 20001
	Error_token_timeout = 20002
	Error_token_create  = 20003
	Error_token         = 20004
	Error_token_auth    = 20005
)

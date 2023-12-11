package pkg_constant

var MsgMap = map[int]string{
	Success:              "success",
	Fail:                 "fail",
	Error_invalid_params: "请求参数异常",

	Error_database_CRUD: "数据库CRUD失败",

	Error_user_exist:          "用户已存在",
	Error_user_not_exist:      "用户不存在",
	Error_user_disabled:       "用户被禁用",
	Error_user_roles_disabled: "用户角色被禁用",
	Error_user_password:       "密码错误",
	Error_user_captcha:        "验证码获取失败",
	Error_user_logout:         "退出失败",

	Error_token_check:   "Token鉴权失败",
	Error_token_timeout: "Token已超时",
	Error_token_create:  "Token生成失败",
	Error_token:         "Token错误",
	Error_token_auth:    "无权限，请联系管理员",
}

func StatusMsg(code int) string {
	msg, ok := MsgMap[code]
	if ok {
		return msg
	}

	return ""
}

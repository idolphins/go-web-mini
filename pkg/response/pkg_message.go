package pkg_response

type ResponseMessage struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
	UIMsg  string `json:"ui_msg"`
	Cancel string `json:"cancel"`
	OK     string `json:"ok"`
}

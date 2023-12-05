package internal

type Response struct {
	Code int         `json:"code"`
	Msg  Message     `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponsePage struct {
	Code      int         `json:"code"`
	Msg       Message     `json:"msg"`
	UIMsg     string      `json:"ui_msg"`
	Data      interface{} `json:"data"`
	Count     int         `json:"count"`
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
}

type Message struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
	UIMsg  string `json:"ui_msg"`
	Cancel string `json:"cancel"`
	OK     string `json:"ok"`
}

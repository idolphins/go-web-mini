package internal

type Response struct {
	Code int `json:"code"`
	Msg  Message
	Data interface{} `json:"data"`
}

type ResponsePage struct {
	Code      int `json:"code"`
	Msg       Message
	Data      interface{} `json:"data"`
	Count     int         `json:"count"`
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
}

type Message struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
	Cancel string `json:"cancel"`
	OK     string `json:"ok"`
}

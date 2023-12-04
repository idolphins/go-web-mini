package pkg_response

import (
	"net/http"
	"osstp-go-hive/pkg/response/internal"

	"github.com/gin-gonic/gin"
)

type Ctx struct {
	Context *gin.Context
}

// success
func Success(c *gin.Context, data interface{}, message string) {
	ctx := Ctx{Context: c}
	messages := ResponseMessage{Msg: message}
	// 200
	ctx.Response(http.StatusOK, data, messages)
}

// fail
func Fail(c *gin.Context, data interface{}, message string) {
	ctx := Ctx{Context: c}
	messages := ResponseMessage{Msg: message}
	// 400
	ctx.Response(http.StatusBadRequest, data, messages)
}

// custom response
func (ctx *Ctx) Response(httpCode int, data interface{}, message ResponseMessage) {
	var (
		ok string = ""
	)
	// code: fail=0, success=1
	code, msg := internal.ResponseCode(httpCode, message.Code, message.Msg)
	if message.OK == "" {
		ok = "ok"
	}
	ctx.Context.JSON(httpCode, internal.ResponsePage{
		Code: code,
		Msg:  internal.Message{Code: message.Code, Title: message.Title, Msg: msg, Cancel: message.Cancel, OK: ok},
		Data: data,
	})
}

// custom response with page and index
func (ctx *Ctx) ResponsePage(httpCode int, data interface{}, count, pageIndex, pageSize int, message ResponseMessage) {
	var (
		ok string = ""
	)
	// code: fail=0, success=1
	// msg: default
	code, msg := internal.ResponseCode(httpCode, message.Code, message.Msg)
	if message.OK == "" {
		ok = "ok"
	}
	ctx.Context.JSON(httpCode, internal.ResponsePage{
		Code:      code,
		Msg:       internal.Message{Code: message.Code, Title: message.Title, Msg: msg, Cancel: message.Cancel, OK: ok},
		Data:      data,
		Count:     count,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	})
}

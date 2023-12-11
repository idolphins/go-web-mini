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
	messages := ResponseMessage{Code: 1, Msg: message}
	// 200
	ctx.Response(http.StatusOK, data, messages)
}

// fail
func Fail(c *gin.Context, data interface{}, message string) {
	ctx := Ctx{Context: c}
	messages := ResponseMessage{Code: 0, Msg: message, OK: "OK"}
	// 400
	ctx.Response(http.StatusBadRequest, data, messages)
}

// custom response
func (ctx *Ctx) Response(httpCode int, data interface{}, message ResponseMessage) {
	// code: fail=0, success=1
	code, msgCode := internal.ResponseCode(httpCode, message.Code)
	ctx.Context.JSON(httpCode, internal.Response{
		Code: code,
		Msg:  internal.Message{Code: msgCode, Title: message.Title, Msg: message.Msg, Cancel: message.Cancel, OK: internal.GetOk(message.OK)},
		Data: data,
	})
}

// custom response with page and index
func (ctx *Ctx) ResponsePage(httpCode int, data interface{}, count, pageIndex, pageSize int, message ResponseMessage) {
	// code: fail=0, success=1
	// msg: default
	code, msgCode := internal.ResponseCode(httpCode, message.Code)
	ctx.Context.JSON(httpCode, internal.ResponsePage{
		Code:      code,
		Msg:       internal.Message{Code: msgCode, Title: message.Title, Msg: message.Msg, Cancel: message.Cancel, OK: internal.GetOk(message.OK)},
		Data:      data,
		Count:     count,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	})
}

package pkg_response

import (
	"go-web-mini/pkg/response/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ctx struct {
	Context *gin.Context
}

// 返回前端
// func Response(c *gin.Context, httpStatus int, code int, data gin.H, message string) {
// 	c.JSON(httpStatus, gin.H{"code": code, "data": data, "message": message})
// }

// success
func Success(c *gin.Context, data interface{}, message string) {
	ctx := Ctx{Context: c}
	messages := ResponseMessage{Msg: message}
	// 200
	ctx.Response(http.StatusOK, 1, data, messages)
}

// fail
func Fail(c *gin.Context, data interface{}, message string) {
	ctx := Ctx{Context: c}
	messages := ResponseMessage{Msg: message}
	// 400
	ctx.Response(http.StatusBadRequest, 0, data, messages)
}

// custom response
func (ctx *Ctx) Response(httpCode int, pkgCode interface{}, data interface{}, message ResponseMessage) {
	code := internal.ResponseCode(httpCode)
	ctx.Context.JSON(httpCode, internal.Response{
		Code: code,
		Msg:  internal.Message{Code: message.Code, Title: message.Title, Msg: message.Msg, Cancel: message.Cancel, OK: message.OK},
		Data: data,
	})
	return
}

// custom response with page and index
func (ctx *Ctx) ResponsePage(httpCode int, pkgCode interface{}, data interface{}, count, pageIndex, pageSize int, message ResponseMessage) {
	code := internal.ResponseCode(httpCode)
	ctx.Context.JSON(httpCode, internal.ResponsePage{
		Code:      code,
		Msg:       internal.Message{Code: message.Code, Title: message.Title, Msg: message.Msg, Cancel: message.Cancel, OK: message.OK},
		Data:      data,
		Count:     count,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	})
	return
}

package app

import (
	"net/http"
	"os"

	"go-quick-template/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	hostname, _ := os.Hostname()
	if data == nil {
		data = gin.H{
			"code":      0,
			"msg":       "success",
			"tracehost": hostname,
		}
	} else {
		data = gin.H{
			"code":      0,
			"msg":       "success",
			"data":      data,
			"tracehost": hostname,
		}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}

package server

import (
	"github.com/gin-gonic/gin"
)

// NewHttpContext 基于 gin.Context 创建一个新的 HttpContext
func NewHttpContext(ctx *gin.Context) *HttpContext {
	hc := &HttpContext{
		Context: ctx,
	}
	return hc
}

// HttpContext 基于 gin.Context 的 http 请求上下文
type HttpContext struct {
	*gin.Context
}

// Gin 获取 gin.Context
func (slf *HttpContext) Gin() *gin.Context {
	return slf.Context
}

// ReadTo 读取请求数据到指定结构体，如果失败则返回错误
func (slf *HttpContext) ReadTo(dest any) error {
	var ctx = slf.Gin()
	if ctx == nil {
		return nil
	}
	if err := ctx.ShouldBind(dest); err != nil {
		if uri := ctx.ShouldBindUri(dest); uri == nil {
			return uri
		} else if query := ctx.ShouldBindQuery(dest); query == nil {
			return query
		}
		return err
	}
	return nil
}

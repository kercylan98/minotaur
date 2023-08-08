package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HttpWrapperGroup http 包装器
type HttpWrapperGroup[CTX any] struct {
	wrapper *HttpWrapper[CTX]
	group   *gin.RouterGroup
}

// Handle 处理请求
func (slf *HttpWrapperGroup[CTX]) Handle(httpMethod, relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	slf.group.Handle(httpMethod, relativePath, handlersToGinHandlers(slf.wrapper.packHandle, handlers)...)
	return slf
}

// Use 使用中间件
func (slf *HttpWrapperGroup[CTX]) Use(middleware ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	slf.group.Use(handlersToGinHandlers(slf.wrapper.packHandle, middleware)...)
	return slf
}

// GET 注册 GET 请求
func (slf *HttpWrapperGroup[CTX]) GET(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return slf.Handle(http.MethodGet, relativePath, handlers...)
}

// POST 注册 POST 请求
func (slf *HttpWrapperGroup[CTX]) POST(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return slf.Handle(http.MethodPost, relativePath, handlers...)
}

// DELETE 注册 DELETE 请求
func (slf *HttpWrapperGroup[CTX]) DELETE(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return slf.Handle(http.MethodDelete, relativePath, handlers...)
}

// PATCH 注册 PATCH 请求
func (slf *HttpWrapperGroup[CTX]) PATCH(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return slf.Handle(http.MethodPatch, relativePath, handlers...)
}

// PUT 注册 PUT 请求
func (slf *HttpWrapperGroup[CTX]) PUT(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return slf.Handle(http.MethodPut, relativePath, handlers...)
}

// OPTIONS 注册 OPTIONS 请求
func (slf *HttpWrapperGroup[CTX]) OPTIONS(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return slf.Handle(http.MethodOptions, relativePath, handlers...)
}

// Group 创建分组
func (slf *HttpWrapperGroup[CTX]) Group(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return &HttpWrapperGroup[CTX]{
		wrapper: slf.wrapper,
		group:   slf.wrapper.server.Group(relativePath, handlersToGinHandlers(slf.wrapper.packHandle, handlers)...),
	}
}

package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpWrapperHandleFunc[CTX any] func(ctx CTX)

func NewHttpWrapper[CTX any](server *Server, pack func(ctx *gin.Context) CTX) *HttpWrapper[CTX] {
	return &HttpWrapper[CTX]{
		server:     server.HttpRouter().(*gin.Engine),
		packHandle: pack,
	}
}

// HttpWrapper http 包装器
type HttpWrapper[CTX any] struct {
	server     *gin.Engine
	packHandle func(ctx *gin.Context) CTX
}

// handlersToGinHandlers 将 HttpWrapperHandleFunc 转换为 gin.HandlerFunc
func handlersToGinHandlers[CTX any](packHandle func(ctx *gin.Context) CTX, handlers []HttpWrapperHandleFunc[CTX]) []gin.HandlerFunc {
	handles := make([]gin.HandlerFunc, len(handlers))
	for i, handle := range handlers {
		handles[i] = func(ctx *gin.Context) {
			handle(packHandle(ctx))
		}
	}
	return handles
}

// Handle 处理请求
func (slf *HttpWrapper[CTX]) Handle(httpMethod, relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	slf.server.Handle(httpMethod, relativePath, handlersToGinHandlers(slf.packHandle, handlers)...)
	return slf
}

// Use 使用中间件
func (slf *HttpWrapper[CTX]) Use(middleware ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	slf.server.Use(handlersToGinHandlers(slf.packHandle, middleware)...)
	return slf
}

// GET 注册 GET 请求
func (slf *HttpWrapper[CTX]) GET(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodGet, relativePath, handlers...)
}

// POST 注册 POST 请求
func (slf *HttpWrapper[CTX]) POST(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodPost, relativePath, handlers...)
}

// DELETE 注册 DELETE 请求
func (slf *HttpWrapper[CTX]) DELETE(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodDelete, relativePath, handlers...)
}

// PATCH 注册 PATCH 请求
func (slf *HttpWrapper[CTX]) PATCH(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodPatch, relativePath, handlers...)
}

// PUT 注册 PUT 请求
func (slf *HttpWrapper[CTX]) PUT(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodPut, relativePath, handlers...)
}

// OPTIONS 注册 OPTIONS 请求
func (slf *HttpWrapper[CTX]) OPTIONS(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodOptions, relativePath, handlers...)
}

// HEAD 注册 HEAD 请求
func (slf *HttpWrapper[CTX]) HEAD(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodHead, relativePath, handlers...)
}

// Trace 注册 Trace 请求
func (slf *HttpWrapper[CTX]) Trace(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodTrace, relativePath, handlers...)
}

// Connect 注册 Connect 请求
func (slf *HttpWrapper[CTX]) Connect(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	return slf.Handle(http.MethodConnect, relativePath, handlers...)
}

// Any 注册 Any 请求
func (slf *HttpWrapper[CTX]) Any(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	slf.Handle(http.MethodGet, relativePath, handlers...)
	slf.Handle(http.MethodPost, relativePath, handlers...)
	slf.Handle(http.MethodDelete, relativePath, handlers...)
	slf.Handle(http.MethodPatch, relativePath, handlers...)
	slf.Handle(http.MethodPut, relativePath, handlers...)
	slf.Handle(http.MethodOptions, relativePath, handlers...)
	slf.Handle(http.MethodHead, relativePath, handlers...)
	slf.Handle(http.MethodTrace, relativePath, handlers...)
	slf.Handle(http.MethodConnect, relativePath, handlers...)
	return slf
}

// Match 注册与您声明的指定方法相匹配的路由。
func (slf *HttpWrapper[CTX]) Match(methods []string, relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapper[CTX] {
	slf.server.Match(methods, relativePath, handlersToGinHandlers(slf.packHandle, handlers)...)
	return slf
}

// StaticFile 注册 StaticFile 请求
func (slf *HttpWrapper[CTX]) StaticFile(relativePath, filepath string) *HttpWrapper[CTX] {
	slf.server.StaticFile(relativePath, filepath)
	return slf
}

// Static 注册 Static 请求
func (slf *HttpWrapper[CTX]) Static(relativePath, root string) *HttpWrapper[CTX] {
	slf.server.Static(relativePath, root)
	return slf
}

// StaticFS 注册 StaticFS 请求
func (slf *HttpWrapper[CTX]) StaticFS(relativePath string, fs http.FileSystem) *HttpWrapper[CTX] {
	slf.server.StaticFS(relativePath, fs)
	return slf
}

// Group 创建一个新的路由组。您应该添加所有具有共同中间件的路由。
func (slf *HttpWrapper[CTX]) Group(relativePath string, handlers ...HttpWrapperHandleFunc[CTX]) *HttpWrapperGroup[CTX] {
	return &HttpWrapperGroup[CTX]{
		wrapper: slf,
		group:   slf.server.Group(relativePath, handlersToGinHandlers(slf.packHandle, handlers)...),
	}
}

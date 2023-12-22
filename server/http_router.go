package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HandlerFunc[Context any] func(ctx Context)
type ContextPacker[Context any] func(ctx *gin.Context) Context

type HttpRouter[Context any] struct {
	srv    *Server
	group  gin.IRouter
	packer ContextPacker[Context]
}

func (slf *HttpRouter[Context]) handlesConvert(handlers []HandlerFunc[Context]) []gin.HandlerFunc {
	var handles []gin.HandlerFunc
	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]
		handles = append(handles, func(ctx *gin.Context) {
			slf.srv.hitMessageStatistics()
			defer func() {
				slf.srv.messageCounter.Add(-1)
			}()
			hc := slf.packer(ctx)
			var now = time.Now()
			handler(hc)
			slf.srv.low(nil, now, time.Second, "HTTP ["+ctx.Request.Method+"] "+ctx.Request.RequestURI)
		})
	}
	return handles
}

// Handle 使用给定的路径和方法注册新的请求句柄和中间件
//   - 最后一个处理程序应该是真正的处理程序，其他处理程序应该是可以而且应该在不同路由之间共享的中间件。
func (slf *HttpRouter[Context]) Handle(httpMethod, relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	handles := slf.handlesConvert(handlers)
	slf.group.Handle(httpMethod, relativePath, handles...)
	return slf
}

// POST 是 Handle("POST", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) POST(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodPost, relativePath, handlers...)
}

// GET 是 Handle("GET", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) GET(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodGet, relativePath, handlers...)
}

// DELETE 是 Handle("DELETE", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) DELETE(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodDelete, relativePath, handlers...)
}

// PATCH 是 Handle("PATCH", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) PATCH(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodPatch, relativePath, handlers...)
}

// PUT 是 Handle("PUT", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) PUT(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodPut, relativePath, handlers...)
}

// OPTIONS 是 Handle("OPTIONS", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) OPTIONS(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodOptions, relativePath, handlers...)
}

// HEAD 是 Handle("HEAD", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) HEAD(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodHead, relativePath, handlers...)
}

// CONNECT 是 Handle("CONNECT", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) CONNECT(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodConnect, relativePath, handlers...)
}

// TRACE 是 Handle("TRACE", path, handlers) 的快捷方式
func (slf *HttpRouter[Context]) TRACE(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	return slf.Handle(http.MethodTrace, relativePath, handlers...)
}

// Any 注册一个匹配所有 HTTP 方法的路由
//   - GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (slf *HttpRouter[Context]) Any(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	for _, m := range []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodHead,
		http.MethodOptions, http.MethodDelete, http.MethodConnect, http.MethodTrace} {
		slf.Handle(m, relativePath, handlers...)
	}
	return slf
}

// Match 注册一个匹配指定 HTTP 方法的路由
//   - GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (slf *HttpRouter[Context]) Match(methods []string, relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	for _, m := range methods {
		slf.Handle(m, relativePath, handlers...)
	}
	return slf
}

// StaticFile 注册单个路由以便为本地文件系统的单个文件提供服务。
//   - 例如: StaticFile("favicon.ico", "./resources/favicon.ico")
func (slf *HttpRouter[Context]) StaticFile(relativePath, filepath string) *HttpRouter[Context] {
	slf.group.StaticFile(relativePath, filepath)
	return slf
}

// StaticFileFS 与 `StaticFile` 类似，但可以使用自定义的 `http.FileSystem` 代替。
//   - 例如: StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
//   - 由于依赖于 gin.Engine 默认情况下使用：gin.Dir
func (slf *HttpRouter[Context]) StaticFileFS(relativePath, filepath string, fs http.FileSystem) *HttpRouter[Context] {
	slf.group.StaticFileFS(relativePath, filepath, fs)
	return slf
}

// Static 提供来自给定文件系统根目录的文件。
//   - 例如: Static("/static", "/var/www")
func (slf *HttpRouter[Context]) Static(relativePath, root string) *HttpRouter[Context] {
	slf.group.StaticFS(relativePath, gin.Dir(root, false))
	return slf
}

// StaticFS 与 `Static` 类似，但可以使用自定义的 `http.FileSystem` 代替。
//   - 例如: StaticFS("/static", Dir{"/var/www", false})
//   - 由于依赖于 gin.Engine 默认情况下使用：gin.Dir
func (slf *HttpRouter[Context]) StaticFS(relativePath string, fs http.FileSystem) *HttpRouter[Context] {
	slf.group.StaticFS(relativePath, fs)
	return slf
}

// Group 创建一个新的路由组。您应该添加所有具有共同中间件的路由。
//   - 例如: v1 := slf.Group("/v1")
func (slf *HttpRouter[Context]) Group(relativePath string, handlers ...HandlerFunc[Context]) *HttpRouter[Context] {
	group := slf.group.Group(relativePath, slf.handlesConvert(handlers)...)
	return &HttpRouter[Context]{
		srv:    slf.srv,
		group:  group,
		packer: slf.packer,
	}
}

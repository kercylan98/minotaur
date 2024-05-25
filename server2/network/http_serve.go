package network

import "net/http"

type HttpServe struct {
	*http.ServeMux
}

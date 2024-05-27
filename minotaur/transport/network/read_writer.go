package network

import "io"

type ReadWriter struct {
	io.Reader
	io.Writer
}

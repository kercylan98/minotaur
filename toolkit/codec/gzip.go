package codec

import (
	"bytes"
	"compress/gzip"
	"io"
)

type GZip struct{}

func (g *GZip) Encode(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(src)
	if err != nil {
		return nil, err
	}
	if err = w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g *GZip) Decode(src []byte) ([]byte, error) {
	data := *bytes.NewBuffer(src)
	r, err := gzip.NewReader(&data)
	if err != nil {
		return nil, err
	}
	result, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err = r.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

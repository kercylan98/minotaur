package codec

import (
	"archive/tar"
	"bytes"
	"io"
)

type Tar struct{}

func (t *Tar) Encode(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := tar.NewWriter(&buf)
	hdr := &tar.Header{
		Name: "file",
		Mode: 0600,
		Size: int64(len(src)),
	}
	if err := w.WriteHeader(hdr); err != nil {
		return nil, err
	}
	if _, err := w.Write(src); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *Tar) Decode(src []byte) ([]byte, error) {
	data := bytes.NewReader(src)
	r := tar.NewReader(data)
	var result bytes.Buffer
	for {
		_, err := r.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(&result, r); err != nil {
			return nil, err
		}
	}
	return result.Bytes(), nil
}

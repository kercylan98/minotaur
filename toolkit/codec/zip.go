package codec

import (
	"archive/zip"
	"bytes"
	"io"
)

type Zip struct{}

func (z *Zip) Encode(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	f, err := w.Create("file")
	if err != nil {
		return nil, err
	}
	_, err = f.Write(src)
	if err != nil {
		return nil, err
	}
	if err = w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (z *Zip) Decode(src []byte) ([]byte, error) {
	data := bytes.NewReader(src)
	r, err := zip.NewReader(data, int64(len(src)))
	if err != nil {
		return nil, err
	}
	var result bytes.Buffer
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(&result, rc)
		if err != nil {
			return nil, err
		}
		_ = rc.Close()
	}
	return result.Bytes(), nil
}

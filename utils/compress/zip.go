package compress

import (
	"archive/zip"
	"bytes"
	"io"
)

// ZIPCompress 对数据进行ZIP压缩，返回bytes.Buffer和错误信息
func ZIPCompress(data []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	f, err := zipWriter.Create("file")
	if err != nil {
		return buf, err
	}
	_, err = f.Write(data)
	if err != nil {
		return buf, err
	}
	if err := zipWriter.Close(); err != nil {
		return buf, err
	}
	return buf, nil
}

// ZIPUnCompress 对已进行ZIP压缩的数据进行解压缩，返回字节数组及错误信息
func ZIPUnCompress(dataByte []byte) ([]byte, error) {
	data := bytes.NewReader(dataByte)
	zipReader, err := zip.NewReader(data, int64(len(dataByte)))
	if err != nil {
		return nil, err
	}
	var result bytes.Buffer
	for _, f := range zipReader.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(&result, rc)
		if err != nil {
			return nil, err
		}
		rc.Close()
	}
	return result.Bytes(), nil
}

package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

// GZipCompress 对数据进行GZip压缩，返回bytes.Buffer和错误信息
func GZipCompress(data []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err := gzipWriter.Write(data)
	if err != nil {
		return buf, err
	}
	if err := gzipWriter.Close(); err != nil {
		return buf, err
	}
	return buf, nil
}

// GZipUnCompress 对已进行GZip压缩的数据进行解压缩，返回字节数组及错误信息
func GZipUnCompress(dataByte []byte) ([]byte, error) {
	data := *bytes.NewBuffer(dataByte)
	gzipReader, err := gzip.NewReader(&data)
	if err != nil {
		return nil, err
	}
	result, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	if err := gzipReader.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

package compress

import (
	"archive/tar"
	"bytes"
	"io"
)

// TARCompress 对数据进行TAR压缩，返回bytes.Buffer和错误信息
func TARCompress(data []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer
	tarWriter := tar.NewWriter(&buf)
	hdr := &tar.Header{
		Name: "file",
		Mode: 0600,
		Size: int64(len(data)),
	}
	if err := tarWriter.WriteHeader(hdr); err != nil {
		return buf, err
	}
	if _, err := tarWriter.Write(data); err != nil {
		return buf, err
	}
	if err := tarWriter.Close(); err != nil {
		return buf, err
	}
	return buf, nil
}

// TARUnCompress 对已进行TAR压缩的数据进行解压缩，返回字节数组及错误信息
func TARUnCompress(dataByte []byte) ([]byte, error) {
	data := bytes.NewReader(dataByte)
	tarReader := tar.NewReader(data)
	var result bytes.Buffer
	for {
		_, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(&result, tarReader); err != nil {
			return nil, err
		}
	}
	return result.Bytes(), nil
}

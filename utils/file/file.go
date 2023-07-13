package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

// PathExist 路径是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsDir 路径是否是文件夹
func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.IsDir(), nil
	}
	return false, err
}

// WriterFile 向特定文件写入内容
func WriterFile(filePath string, content []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, io.SeekEnd)
		_, err = f.WriteAt(content, n)
		_ = f.Close()
	}
	return nil
}

// ReadOnce 单次读取文件
//   - 一次性对整个文件进行读取，小文件读取可以很方便的一次性将文件内容读取出来，而大文件读取会造成性能影响。
func ReadOnce(filePath string) ([]byte, error) {
	if file, err := os.Open(filePath); err != nil {
		return nil, err
	} else {
		defer func() {
			_ = file.Close()
		}()
		return io.ReadAll(file)
	}
}

// ReadBlockHook 分块读取文件
//   - 将filePath路径对应的文件数据并将读到的每一部分传入hook函数中，当过程中如果产生错误则会返回error。
//   - 分块读取可以在读取速度和内存消耗之间有一个很好的平衡。
func ReadBlockHook(filePath string, bufferSize int, hook func(data []byte)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	buffer := make([]byte, bufferSize)
	bufferReader := bufio.NewReader(file)
	for {
		successReadSize, err := bufferReader.Read(buffer)
		hook(buffer[:successReadSize])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// LineCount 统计文件行数
func LineCount(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}

	line := 0
	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			line++
		}
	}
	return line
}

// FilePaths 获取指定目录下的所有文件路径
//   - 包括了子目录下的文件
//   - 不包含目录
func FilePaths(dir string) []string {
	var paths []string
	abs, err := filepath.Abs(dir)
	if err != nil {
		return paths
	}
	files, err := os.ReadDir(abs)
	if err != nil {
		return paths
	}

	for _, file := range files {
		fileAbs := filepath.Join(abs, file.Name())
		if file.IsDir() {
			paths = append(paths, FilePaths(fileAbs)...)
			continue
		}
		paths = append(paths, fileAbs)
	}
	return paths
}

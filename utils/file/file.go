package file

import (
	"bufio"
	"github.com/kercylan98/minotaur/utils/slice"
	"io"
	"os"
	"path/filepath"
	"sync"
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

// ReadLine 分行读取文件
//   - 将filePath路径对应的文件数据并将读到的每一行传入hook函数中，当过程中如果产生错误则会返回error。
func ReadLine(filePath string, hook func(line string)) error {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		hook(string(line))
	}
	return nil
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

// Paths 获取指定目录下的所有文件路径
//   - 包括了子目录下的文件
//   - 不包含目录
func Paths(dir string) []string {
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
			paths = append(paths, Paths(fileAbs)...)
			continue
		}
		paths = append(paths, fileAbs)
	}
	return paths
}

// ReadLineWithParallelByChannel 并行的分行读取文件并行处理，处理过程中会将每一行的内容传入 handlerFunc 中进行处理
//   - 由于是并行处理，所以处理过程中的顺序是不确定的。
//   - 可通过 start 参数指定开始读取的位置，如果不指定则从文件开头开始读取。
func ReadLineWithParallel(filename string, chunkSize int64, handlerFunc func(string), start ...int64) (n int64, err error) {
	file, err := os.Open(filename)
	offset := slice.GetValue(start, 0)
	if err != nil {
		return offset, err
	}
	defer func() {
		_ = file.Close()
	}()

	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return offset, err
	}
	if offset-1 >= fileSize {
		return fileSize + 1, nil
	}

	chunks := FindLineChunksByOffset(file, offset, chunkSize)
	var end int64
	var endMutex sync.Mutex
	var wg sync.WaitGroup
	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk [2]int64) {
			defer wg.Done()

			endMutex.Lock()
			e := chunk[1] - chunk[0]
			if e > end {
				end = e + 1
			}
			endMutex.Unlock()
			r := io.NewSectionReader(file, chunk[0], e)

			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				handlerFunc(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				return
			}
		}(chunk)
	}
	wg.Wait()
	return end, nil
}

// FindLineChunks 查找文件按照每行划分的分块，每个分块的大小将在 chunkSize 和分割后的分块距离行首及行尾的距离中范围内
//   - 使用该函数得到的分块是完整的行，不会出现行被分割的情况
//   - 当过程中发生错误将会发生 panic
//   - 返回值的成员是一个长度为 2 的数组，第一个元素是分块的起始位置，第二个元素是分块的结束位置
func FindLineChunks(file *os.File, chunkSize int64) [][2]int64 {
	return FindLineChunksByOffset(file, 0, chunkSize)
}

// FindLineChunksByOffset 该函数与 FindLineChunks 类似，不同的是该函数可以指定 offset 从指定位置开始读取文件
func FindLineChunksByOffset(file *os.File, offset, chunkSize int64) [][2]int64 {
	var chunks [][2]int64

	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	currentPos := offset
	for currentPos < fileSize {
		start := currentPos
		if start != 0 { // 不是文件的开头
			for {
				b := make([]byte, 1)
				if _, err = file.ReadAt(b, start); err != nil {
					panic(err)
				}
				if b[0] == '\n' {
					start++ // 移动到下一行的开始
					break
				}
				start--
			}
		}

		end := start + chunkSize
		if end < fileSize { // 不是文件的末尾
			for {
				b := make([]byte, 1)
				if _, err = file.ReadAt(b, end); err != nil {
					panic(err)
				}
				if b[0] == '\n' {
					break
				}
				end++
			}
		} else {
			end = fileSize
		}

		chunks = append(chunks, [2]int64{start, end})
		currentPos = end + 1
	}

	return chunks
}

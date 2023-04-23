package str

import "strings"

// RangeLine 对传入的eachString进行按行切片后再进行遍历
//   - 该函数会预先对“\r\n”进行处理替换为“\n”。
//   - 在遍历到每一行的时候会将结果index和line作为入参传入eachFunc中进行调用。
//   - index表示了当前行的行号（由0开始），line表示了当前行的内容。
func RangeLine(eachString string, eachFunc func(index int, line string) error) error {
	formatStr := strings.ReplaceAll(eachString, "\r\n", "\n")
	for index, line := range strings.Split(formatStr, "\n") {
		if err := eachFunc(index, line); err != nil {
			return err
		}
	}
	return nil
}

package survey

import "github.com/tidwall/gjson"

type (
	Result = gjson.Result
)

// R 记录器所记录的一条数据
type R string

// Get 获取指定 key 的值
//   - 当 key 为嵌套 key 时，使用 . 进行分割，例如：a.b.c
//   - 更多用法参考：https://github.com/tidwall/gjson
func (slf R) Get(key string) Result {
	return gjson.Get(string(slf), key)
}

func (slf R) String() string {
	return string(slf)
}

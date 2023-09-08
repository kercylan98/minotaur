package survey

import (
	"github.com/kercylan98/minotaur/utils/times"
	"github.com/tidwall/gjson"
	"time"
)

type (
	Result = gjson.Result
)

// R 记录器所记录的一条数据
type R string

// GetTime 获取该记录的时间
func (slf R) GetTime(layout string) time.Time {
	return times.GetTimeFromString(string(slf)[:len(layout)], layout)
}

// Get 获取指定 key 的值
//   - 当 key 为嵌套 key 时，使用 . 进行分割，例如：a.b.c
//   - 更多用法参考：https://github.com/tidwall/gjson
func (slf R) Get(key string) Result {
	return gjson.Get(string(slf), key)
}

// Exist 判断指定 key 是否存在
func (slf R) Exist(key string) bool {
	v := slf.Get(key)
	return v.Exists() && len(v.String()) > 0
}

// GetString 该函数为 Get(key).String() 的简写
func (slf R) GetString(key string) string {
	return slf.Get(key).String()
}

// GetInt64 该函数为 Get(key).Int() 的简写
func (slf R) GetInt64(key string) int64 {
	return slf.Get(key).Int()
}

// GetInt 该函数为 Get(key).Int() 的简写，但是返回值为 int 类型
func (slf R) GetInt(key string) int {
	return int(slf.Get(key).Int())
}

// GetFloat64 该函数为 Get(key).Float() 的简写
func (slf R) GetFloat64(key string) float64 {
	return slf.Get(key).Float()
}

// GetBool 该函数为 Get(key).Bool() 的简写
func (slf R) GetBool(key string) bool {
	return slf.Get(key).Bool()
}

func (slf R) String() string {
	return string(slf)
}

package sole

import "github.com/sony/sonyflake"

var sonyflakeGenerator *sonyflake.Sonyflake

func init() {
	sonyflakeGenerator = sonyflake.NewSonyflake(sonyflake.Settings{})
}

// SonyflakeIDE 获取一个雪花id
func SonyflakeIDE() (int64, error) {
	id, err := sonyflakeGenerator.NextID()
	return int64(id), err
}

// SonyflakeID 获取一个雪花id
func SonyflakeID() int64 {
	id, err := sonyflakeGenerator.NextID()
	if err != nil {
		panic(err)
	}
	return int64(id)
}

// SonyflakeSetting 配置雪花id生成策略
func SonyflakeSetting(settings sonyflake.Settings) {
	sonyflakeGenerator = sonyflake.NewSonyflake(settings)
}

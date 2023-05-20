package network

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/maths"
	"strings"
)

// FloatEnlarge 用于将浮点型放大后进行网络传输，返回放大后的值和放大的倍率
//   - 存在精度丢失问题，如1.13
func FloatEnlarge[F generic.Float](f F) (value int64, multi int64) {
	str := fmt.Sprint(f)
	multi = maths.PowInt64(10, int64(len(str[strings.Index(str, ".")+1:])))
	value = int64(f * F(multi))
	return
}

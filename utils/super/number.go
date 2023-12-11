package super

import "reflect"

var (
	romeThousands = []string{"", "M", "MM", "MMM"}
	romeHundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	romeTens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	romeOnes      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

// IsNumber 判断是否为数字
func IsNumber(v any) bool {
	kind := reflect.Indirect(reflect.ValueOf(v)).Kind()
	return kind >= reflect.Int && kind <= reflect.Float64
}

// NumberToRome 将数字转换为罗马数字
func NumberToRome(num int) string {
	return romeThousands[num/1000] + romeHundreds[num%1000/100] + romeTens[num%100/10] + romeOnes[num%10]
}

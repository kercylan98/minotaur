package huge_test

import (
	"github.com/kercylan98/minotaur/utils/huge"
	"math/big"
	"testing"
)

func TestNewInt(t *testing.T) {
	var cases = []struct {
		name string
		nil  bool
		in   int64
		mul  int64
		want string
	}{
		{name: "TestNewIntNegative", in: -1, want: "-1"},
		{name: "TestNewIntZero", in: 0, want: "0"},
		{name: "TestNewIntPositive", in: 1, want: "1"},
		{name: "TestNewIntMax", in: 9223372036854775807, want: "9223372036854775807"},
		{name: "TestNewIntMin", in: -9223372036854775808, want: "-9223372036854775808"},
		{name: "TestNewIntMulNegative", in: -9223372036854775808, mul: 10000000, want: "-92233720368547758080000000"},
		{name: "TestNewIntMulPositive", in: 9223372036854775807, mul: 10000000, want: "92233720368547758070000000"},
		{name: "TestNewIntNil", nil: true, want: "0"},
		{name: "TestNewIntNilMul", nil: true, mul: 10000000, want: "0"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got *huge.Int
			switch {
			case c.nil:
				if c.mul > 0 {
					got = huge.NewInt(0).MulInt64(c.mul)
				}
			case c.mul == 0:
				got = huge.NewInt(c.in)
			default:
				got = huge.NewInt(c.in).MulInt64(c.mul)
			}
			if s := got.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, got.String())
			} else {
				t.Log(s)
			}
		})
	}

	// other
	t.Run("TestNewIntFromString", func(t *testing.T) {
		if got := huge.NewInt("1234567890123456789012345678901234567890"); got.String() != "1234567890123456789012345678901234567890" {
			t.Fatalf("want: %s, got: %s", "1234567890123456789012345678901234567890", got.String())
		}
	})
	t.Run("TestNewIntFromInt", func(t *testing.T) {
		if got := huge.NewInt(1234567890); got.String() != "1234567890" {
			t.Fatalf("want: %s, got: %s", "1234567890", got.String())
		}
	})
	t.Run("TestNewIntFromBool", func(t *testing.T) {
		if got := huge.NewInt(true); got.String() != "1" {
			t.Fatalf("want: %s, got: %s", "1", got.String())
		}
	})
	t.Run("TestNewIntFromFloat", func(t *testing.T) {
		if got := huge.NewInt(1234567890.1234567890); got.String() != "1234567890" {
			t.Fatalf("want: %s, got: %s", "1234567890", got.String())
		}
	})
}

func TestInt_Copy(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{
		{name: "TestIntCopyNegative", in: -1, want: "-1"},
		{name: "TestIntCopyZero", in: 0, want: "0"},
		{name: "TestIntCopyPositive", in: 1, want: "1"},
		{name: "TestIntCopyMax", in: 9223372036854775807, want: "9223372036854775807"},
		{name: "TestIntCopyMin", in: -9223372036854775808, want: "-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in = huge.NewInt(c.in)
			var got = in.Copy()
			if in.Int64() != c.in {
				t.Fatalf("want: %d, got: %d", c.in, in.Int64())
			}
			if s := got.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, got.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_Set(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{
		{name: "TestIntSetNegative", in: -1, want: "-1"},
		{name: "TestIntSetZero", in: 0, want: "0"},
		{name: "TestIntSetPositive", in: 1, want: "1"},
		{name: "TestIntSetMax", in: 9223372036854775807, want: "9223372036854775807"},
		{name: "TestIntSetMin", in: -9223372036854775808, want: "-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.Set(huge.NewInt(c.in))
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetString(t *testing.T) {
	var cases = []struct {
		name string
		in   string
		want string
	}{
		{name: "TestIntSetStringNegative", in: "-1", want: "-1"},
		{name: "TestIntSetStringZero", in: "0", want: "0"},
		{name: "TestIntSetStringPositive", in: "1", want: "1"},
		{name: "TestIntSetStringMax", in: "9223372036854775807", want: "9223372036854775807"},
		{name: "TestIntSetStringMin", in: "-9223372036854775808", want: "-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetString(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetInt(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{
		{name: "TestIntSetIntNegative", in: -1, want: "-1"},
		{name: "TestIntSetIntZero", in: 0, want: "0"},
		{name: "TestIntSetIntPositive", in: 1, want: "1"},
		{name: "TestIntSetIntMax", in: 9223372036854775807, want: "9223372036854775807"},
		{name: "TestIntSetIntMin", in: -9223372036854775808, want: "-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetInt8(t *testing.T) {
	var cases = []struct {
		name string
		in   int8
		want string
	}{
		{name: "TestIntSetInt8Negative", in: -1, want: "-1"},
		{name: "TestIntSetInt8Zero", in: 0, want: "0"},
		{name: "TestIntSetInt8Positive", in: 1, want: "1"},
		{name: "TestIntSetInt8Max", in: 127, want: "127"},
		{name: "TestIntSetInt8Min", in: -128, want: "-128"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt8(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetInt16(t *testing.T) {
	var cases = []struct {
		name string
		in   int16
		want string
	}{
		{name: "TestIntSetInt16Negative", in: -1, want: "-1"},
		{name: "TestIntSetInt16Zero", in: 0, want: "0"},
		{name: "TestIntSetInt16Positive", in: 1, want: "1"},
		{name: "TestIntSetInt16Max", in: 32767, want: "32767"},
		{name: "TestIntSetInt16Min", in: -32768, want: "-32768"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt16(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetInt32(t *testing.T) {
	var cases = []struct {
		name string
		in   int32
		want string
	}{
		{name: "TestIntSetInt32Negative", in: -1, want: "-1"},
		{name: "TestIntSetInt32Zero", in: 0, want: "0"},
		{name: "TestIntSetInt32Positive", in: 1, want: "1"},
		{name: "TestIntSetInt32Max", in: 2147483647, want: "2147483647"},
		{name: "TestIntSetInt32Min", in: -2147483648, want: "-2147483648"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt32(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetInt64(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want string
	}{
		{name: "TestIntSetInt64Negative", in: -1, want: "-1"},
		{name: "TestIntSetInt64Zero", in: 0, want: "0"},
		{name: "TestIntSetInt64Positive", in: 1, want: "1"},
		{name: "TestIntSetInt64Max", in: 9223372036854775807, want: "9223372036854775807"},
		{name: "TestIntSetInt64Min", in: -9223372036854775808, want: "-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetInt64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetUint(t *testing.T) {
	var cases = []struct {
		name string
		in   uint64
		want string
	}{
		{name: "TestIntSetUintNegative", in: 0, want: "0"},
		{name: "TestIntSetUintZero", in: 0, want: "0"},
		{name: "TestIntSetUintPositive", in: 1, want: "1"},
		{name: "TestIntSetUintMax", in: 18446744073709551615, want: "18446744073709551615"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetUint8(t *testing.T) {
	var cases = []struct {
		name string
		in   uint8
		want string
	}{
		{name: "TestIntSetUint8Negative", in: 0, want: "0"},
		{name: "TestIntSetUint8Zero", in: 0, want: "0"},
		{name: "TestIntSetUint8Positive", in: 1, want: "1"},
		{name: "TestIntSetUint8Max", in: 255, want: "255"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint8(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetUint16(t *testing.T) {
	var cases = []struct {
		name string
		in   uint16
		want string
	}{
		{name: "TestIntSetUint16Negative", in: 0, want: "0"},
		{name: "TestIntSetUint16Zero", in: 0, want: "0"},
		{name: "TestIntSetUint16Positive", in: 1, want: "1"},
		{name: "TestIntSetUint16Max", in: 65535, want: "65535"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint16(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetUint32(t *testing.T) {
	var cases = []struct {
		name string
		in   uint32
		want string
	}{
		{name: "TestIntSetUint32Negative", in: 0, want: "0"},
		{name: "TestIntSetUint32Zero", in: 0, want: "0"},
		{name: "TestIntSetUint32Positive", in: 1, want: "1"},
		{name: "TestIntSetUint32Max", in: 4294967295, want: "4294967295"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint32(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetUint64(t *testing.T) {
	var cases = []struct {
		name string
		in   uint64
		want string
	}{
		{name: "TestIntSetUint64Negative", in: 0, want: "0"},
		{name: "TestIntSetUint64Zero", in: 0, want: "0"},
		{name: "TestIntSetUint64Positive", in: 1, want: "1"},
		{name: "TestIntSetUint64Max", in: 18446744073709551615, want: "18446744073709551615"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetUint64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetFloat32(t *testing.T) {
	var cases = []struct {
		name string
		in   float32
		want string
	}{
		{name: "TestIntSetFloat32Negative", in: -1.1, want: "-1"},
		{name: "TestIntSetFloat32Zero", in: 0, want: "0"},
		{name: "TestIntSetFloat32Positive", in: 1.1, want: "1"},
		{name: "TestIntSetFloat32Max", in: 9223372036854775807, want: "9223372036854775807"},
		{name: "TestIntSetFloat32Min", in: -9223372036854775808, want: "-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetFloat32(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetFloat64(t *testing.T) {
	var cases = []struct {
		name string
		in   float64
		want string
	}{
		{name: "TestIntSetFloat64Negative", in: -1.1, want: "-1"},
		{name: "TestIntSetFloat64Zero", in: 0, want: "0"},
		{name: "TestIntSetFloat64Positive", in: 1.1, want: "1"},
		{name: "TestIntSetFloat64Max", in: 9223372036854775807, want: "9223372036854775807"},
		{name: "TestIntSetFloat64Min", in: -9223372036854775808, want: "-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetFloat64(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_SetBool(t *testing.T) {
	var cases = []struct {
		name string
		in   bool
		want string
	}{
		{name: "TestIntSetBoolFalse", in: false, want: "0"},
		{name: "TestIntSetBoolTrue", in: true, want: "1"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var in *huge.Int
			in = in.SetBool(c.in)
			if s := in.String(); s != c.want {
				t.Fatalf("want: %s, got: %s", c.want, in.String())
			} else {
				t.Log(s)
			}
		})
	}
}

func TestInt_IsZero(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want bool
	}{
		{name: "TestIntIsZeroNegative", in: -1, want: false},
		{name: "TestIntIsZeroZero", in: 0, want: true},
		{name: "TestIntIsZeroPositive", in: 1, want: false},
		{name: "TestIntIsZeroMax", in: 9223372036854775807, want: false},
		{name: "TestIntIsZeroMin", in: -9223372036854775808, want: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := huge.NewInt(c.in).IsZero(); got != c.want {
				t.Fatalf("want: %t, got: %t", c.want, got)
			}
		})
	}
}

func TestInt_ToBigint(t *testing.T) {
	var cases = []struct {
		name string
		in   int64
		want *big.Int
	}{
		{name: "TestIntToBigintNegative", in: -1, want: big.NewInt(-1)},
		{name: "TestIntToBigintZero", in: 0, want: big.NewInt(0)},
		{name: "TestIntToBigintPositive", in: 1, want: big.NewInt(1)},
		{name: "TestIntToBigintMax", in: 9223372036854775807, want: big.NewInt(9223372036854775807)},
		{name: "TestIntToBigintMin", in: -9223372036854775808, want: big.NewInt(-9223372036854775808)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := huge.NewInt(c.in).ToBigint(); got.Cmp(c.want) != 0 {
				t.Fatalf("want: %s, got: %s", c.want.String(), got.String())
			}
		})
	}
}

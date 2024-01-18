package huge_test

import (
	"github.com/kercylan98/minotaur/utils/huge"
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
				t.Errorf("want: %s, got: %s", c.want, got.String())
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

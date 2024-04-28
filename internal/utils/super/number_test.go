package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestNumberToRome(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{
		{input: 1, output: "I"},
		{input: 5, output: "V"},
		{input: 10, output: "X"},
		{input: 50, output: "L"},
		{input: 100, output: "C"},
		{input: 500, output: "D"},
		{input: 1000, output: "M"},
	}

	for _, test := range tests {
		result := super.NumberToRome(test.input)
		if result != test.output {
			t.Errorf("NumberToRome(%d) = %s; want %s", test.input, result, test.output)
		}
	}
}

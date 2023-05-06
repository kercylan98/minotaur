package analyzer

import "testing"

func TestDefault_Analyze(t *testing.T) {
	var d = new(Default)
	d.Analyze("./template.xlsx")
}

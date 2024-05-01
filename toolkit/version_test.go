package toolkit_test

import (
	"github.com/kercylan98/minotaur/toolkit"
	"testing"
)

func TestOldVersion(t *testing.T) {
	testCases := []struct {
		version1 string
		version2 string
		want     bool
	}{
		{"1.2.3", "1.2.2", true},
		{"1.2.1", "1.2.2", false},
		{"1.2.3", "1.2.3", false},
		{"v1.2.3", "v1.2.2", true},
		{"v1.2.3", "v1.2.4", false},
		{"v1.2.3", "1.2.3", false},
		{"vxx2faf.d2ad5.dd3", "gga2faf.d2ad5.dd2", true},
		{"awd2faf.d2ad4.dd3", "vsd2faf.d2ad5.dd3", false},
		{"vxd2faf.d2ad5.dd3", "qdq2faf.d2ad5.dd3", false},
		{"1.2.3", "vdafe2faf.d2ad5.dd3", false},
		{"v1.2.3", "vdafe2faf.d2ad5.dd3", false},
	}

	for _, tc := range testCases {
		got := toolkit.OldVersion(tc.version1, tc.version2)
		if got != tc.want {
			t.Errorf("OldVersion(%q, %q) = %v; want %v", tc.version1, tc.version2, got, tc.want)
		}
	}
}

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		version1 string
		version2 string
		want     int
	}{
		{"1.2.3", "1.2.2", 1},
		{"1.2.1", "1.2.2", -1},
		{"1.2.3", "1.2.3", 0},
		{"v1.2.3", "v1.2.2", 1},
		{"v1.2.3", "v1.2.4", -1},
		{"v1.2.3", "1.2.3", 0},
		{"vde2faf.d2ad5.dd3", "e2faf.d2ad5.dd2", 1},
		{"vde2faf.d2ad4.dd3", "vde2faf.d2ad5.dd3", -1},
		{"vfe2faf.d2ad5.dd3", "ve2faf.d2ad5.dd3", 0},
		{"1.2.3", "vdafe2faf.d2ad5.dd3", -1},
		{"v1.2.3", "vdafe2faf.d2ad5.dd3", -1},
	}

	for _, tc := range testCases {
		got := toolkit.CompareVersion(tc.version1, tc.version2)
		if got != tc.want {
			t.Errorf("CompareVersion(%q, %q) = %v; want %v", tc.version1, tc.version2, got, tc.want)
		}
	}
}

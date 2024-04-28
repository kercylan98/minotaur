package super

import (
	"regexp"
	"strings"
)

// OldVersion 检查 version2 对于 version1 来说是不是旧版本
func OldVersion(version1, version2 string) bool {
	return CompareVersion(version1, version2) == 1
}

// CompareVersion 返回一个整数，用于表示两个版本号的比较结果：
//   - 如果 version1 大于 version2，它将返回 1
//   - 如果 version1 小于 version2，它将返回 -1
//   - 如果 version1 和 version2 相等，它将返回 0
func CompareVersion(version1, version2 string) int {
	reg, _ := regexp.Compile("[^0-9.]+")
	processedVersion1 := processVersion(reg.ReplaceAllString(version1, ""))
	processedVersion2 := processVersion(reg.ReplaceAllString(version2, ""))

	n, m := len(processedVersion1), len(processedVersion2)
	i, j := 0, 0
	for i < n || j < m {
		x := 0
		for ; i < n && processedVersion1[i] != '.'; i++ {
			x = x*10 + int(processedVersion1[i]-'0')
		}
		i++ // skip the dot
		y := 0
		for ; j < m && processedVersion2[j] != '.'; j++ {
			y = y*10 + int(processedVersion2[j]-'0')
		}
		j++ // skip the dot
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

func processVersion(version string) string {
	// 移除首尾可能存在的非数字字符
	reg, _ := regexp.Compile("^[^.0-9]+|[^.0-9]+$")
	version = reg.ReplaceAllString(version, "")
	// 确保不出现连续的点
	version = strings.ReplaceAll(version, "..", ".")
	// 移除每一部分起始的 0
	versionParts := strings.Split(version, ".")
	for idx, part := range versionParts {
		versionParts[idx] = strings.TrimLeft(part, "0")
	}
	return strings.Join(versionParts, ".")
}

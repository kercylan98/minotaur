package charproc

import (
	"regexp"
	"strings"
)

// SensitiveTrieNode 敏感词树节点
type SensitiveTrieNode struct {
	nodes map[rune]*SensitiveTrieNode
	match bool
}

// Add 添加敏感词
func (s *SensitiveTrieNode) Add(text string) {
	if s.nodes == nil {
		s.nodes = make(map[rune]*SensitiveTrieNode)
	}
	chars := []rune(strings.ToUpper(text))
	l := len(chars)
	if l == 0 {
		return
	}
	node := s
	for i := 0; i < l; i++ {
		ch := chars[i]
		if _, ok := node.nodes[ch]; !ok {
			node.nodes[ch] = &SensitiveTrieNode{nodes: make(map[rune]*SensitiveTrieNode)}
		}
		node = node.nodes[ch]
	}
	node.match = true
}

// Check 检查是否包含敏感词
func (s *SensitiveTrieNode) Check(chars []rune) bool {
	l := len(chars)
	if l == 0 {
		return false
	}

	nodes := s.nodes
	for i := 0; i < l; i++ {
		ch := s.runeToUpper(chars[i])
		node, ok := nodes[ch]
		if !ok {
			continue
		}
		if node.match {
			return true
		}
		nodes = node.nodes
		for j := i + 1; j < l; j++ {
			ch = s.runeToUpper(chars[j])
			node, ok := nodes[ch]
			if !ok {
				break
			}
			if node.match {
				return true
			}
			nodes = node.nodes
		}
		nodes = s.nodes
	}
	return false
}

// Replace 替换敏感词为指定字符
func (s *SensitiveTrieNode) Replace(chars []rune, rep rune) []rune {
	l := len(chars)
	if l == 0 {
		return chars
	}

	nodes := s.nodes
	for i := 0; i < l; i++ {
		ch := s.runeToUpper(chars[i])
		node, ok := nodes[ch]
		if !ok {
			continue
		}
		if node.match {
			for j := i; j < l; j++ {
				chars[j] = rep
			}
		}
		nodes = node.nodes
		for j := i + 1; j < l; j++ {
			ch = s.runeToUpper(chars[j])
			node, ok := nodes[ch]
			if !ok {
				break
			}
			if node.match {
				for k := i; k <= j; k++ {
					chars[k] = rep
				}
				i = j
				break
			}
			nodes = node.nodes
		}
		nodes = s.nodes
	}
	return chars
}

func (s *SensitiveTrieNode) runeToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r -= 'a' - 'A'
	}
	return r
}

// HideSensitivity 返回防敏感化后的字符串
//   - 隐藏身份证、邮箱、手机号等敏感信息用 * 号替代
func HideSensitivity(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
			return result
		}
		resRs := []rune(str)
		res2 := string(resRs[0:3])
		resString := res2 + "***"
		result = resString + "@" + res[1]
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			rs := []rune(str)
			result = string(rs[0:5]) + "****" + string(rs[7:11])
			return
		}
		nameRune := []rune(str)
		lens := len(nameRune)

		if lens <= 1 {
			result = "***"
		} else if lens == 2 {
			result = string(nameRune[:1]) + "*"
		} else if lens == 3 {
			result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
		} else if lens == 4 {
			result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
		} else {
			result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
		}
		return
	}
}

package stack

import "strings"

// 20. 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/
func isValid1(s string) bool {
	if len(s) == 0 {
		return true
	}
	mapDict := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	runes := []rune(s)
	var stack Stack
	for _, val := range runes {
		if _, ok := mapDict[val]; ok {
			stack.Push(val)
		} else {
			if stack.IsEmpty() {
				return false
			}
			if val != mapDict[stack.Pop().(rune)] {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func isValid2(s string) bool {
	if len(s) == 0 {
		return true
	}
	runes := []rune(s)
	var stack Stack
	for _, val := range runes {
		if val == '(' || val == '{' || val == '[' {
			stack.Push(val)
		} else {
			if stack.IsEmpty() {
				return false
			}
			left := stack.Pop().(rune)
			if left == '(' && val != ')' {
				return false
			}
			if left == '{' && val != '}' {
				return false
			}
			if left == '[' && val != ']' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func isValid3(s string) bool {
	for strings.Contains(s, "()") || strings.Contains(s, "{}") || strings.Contains(s, "[]") {
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
	}
	return len(s) == 0
}

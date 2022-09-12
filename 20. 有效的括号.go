package main

// 获取对应的左括号，简化代码
func getLeftBracket(c rune) rune {
	switch c {
	case ')':
		return '('
	case ']':
		return '['
	default:
		return '{'
	}
}

func isValid(s string) bool {
	var stack []rune // 使用stack来存储左括号
	for _, x := range s {
		n := len(stack)
		switch x {
		case '(', '{', '[':
			stack = append(stack, x)
		default:
			if n <= 0 || stack[n-1] != getLeftBracket(x) { //防止越界需要判断n的大小
				return false
			}
			stack = stack[:n-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

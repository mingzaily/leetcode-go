package main

import "fmt"

func isValid(s string) bool {
	length := len(s)
	// 奇数个
	if length%2 != 0 {
		return false
	}

	m := map[string]string{")": "(", "]": "[", "}": "{"}
	// 栈存放左括号
	var stack []string

	for _, v := range s {
		if value, ok := m[string(v)]; ok {
			// 右括号
			// 如果是右括号，栈却空，表明不匹配
			// 栈顶元素不与右括号的对立面（左括号）相等，表明不匹配
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 左括号入栈
			stack = append(stack, string(v))
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(){}"))
	fmt.Println(isValid("({)}"))
}

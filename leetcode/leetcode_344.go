package main

import "fmt"

func solve(str string) string {
	// write code here
	var temp string
	for i := len(str) - 1; i >= 0; i-- {
		temp += string(str[i])
	}

	return temp
}

func reverseString(s []byte) string {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}

func main() {
	fmt.Println(solve("abcde"))
	fmt.Println(reverseString([]byte("abcde")))
}

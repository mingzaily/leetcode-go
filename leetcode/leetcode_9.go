package main

import "fmt"

func isPalindrome(x int) bool {
	revereNum := 0
	if x < 0 {
		return false
	}

	tempNum := x
	for x != 0 {
		revereNum = revereNum*10 + x%10
		x = x / 10
	}

	if tempNum == revereNum {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
}

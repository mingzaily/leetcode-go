package main

import "fmt"

func firstUniqChar(s string) byte {
	var temp [26]int
	for _, v := range s {
		temp[v-'a']++
	}

	for i, v := range s {
		if temp[v-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
	fmt.Println(firstUniqChar("leetcode"))
}

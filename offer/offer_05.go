package main

import (
	"fmt"
)

// 官方自带
//func replaceSpace(s string) string {
//	count := strings.Count(s, " ")
//	return strings.Replace(s, " ", "%20", count)
//}

func replaceSpace(s string) string {
	str := ""
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

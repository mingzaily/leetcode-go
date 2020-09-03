package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	// 基准字符串
	global := strs[0]
	for _, v := range strs[1:] {
		// 遍历 基准字符串在各个数组元素是否出现
		// 当=0时代表公共前缀已找到或是子串为空字符串，退出循环，进行第二次比较
		for strings.Index(v, global) != 0 {
			// 基准元素长度为0代表无公共前缀
			if len(global) == 0 {
				return ""
			}
			// 循环减少基准字符串的末尾元素直到存在
			global = global[:len(global)-1]
		}
	}
	return global
}

func main() {
	str := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(str))
}

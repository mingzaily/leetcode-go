package main

import "fmt"

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	left, right, max := -1, 0, 0
	temp := map[byte]int{}
	for ; right < len(s); right++ {
		// hash表key为值，value为下标
		// 遇到存在hash表的key，左指针跳到对应重复key的下标，但不能往回跳
		if value, ok := temp[s[right]]; ok {
			left = maxInt(left, value)
		}
		temp[s[right]] = right
		max = maxInt(max, right-left)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbb"))
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))
	//fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("aud"))
	//fmt.Println(lengthOfLongestSubstring("abba"))
}

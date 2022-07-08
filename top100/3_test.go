package leetcode_top100

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	var result int
	hashMap := make(map[int32]int)

	for index, value := range s {
		// 如果存在，滑动窗口左值往右移动
		if lastValueIndex, ok := hashMap[value]; ok {
			// 滑动窗口左值不能往左退
			left = max(lastValueIndex+1, left)
		}
		// 向右增加滑动窗口
		right++
		result = max(right-left, result)
		hashMap[value] = index
		fmt.Println("滑动窗口: ", s[left:right])
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test_LengthOfLongestSubstring1(t *testing.T) {
	s := "pwwkew"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Errorf("Expected: 3, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring2(t *testing.T) {
	s := "a"
	result := lengthOfLongestSubstring(s)
	if result != 1 {
		t.Errorf("Expected: 1, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring3(t *testing.T) {
	s := "dvdf"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Errorf("Expected: 3, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring4(t *testing.T) {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Errorf("Expected: 3, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring5(t *testing.T) {
	s := "aud"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Errorf("Expected: 3, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring6(t *testing.T) {
	s := "dvd"
	result := lengthOfLongestSubstring(s)
	if result != 2 {
		t.Errorf("Expected: 2, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring7(t *testing.T) {
	s := "cdd"
	result := lengthOfLongestSubstring(s)
	if result != 2 {
		t.Errorf("Expected: 2, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring8(t *testing.T) {
	s := "aabc"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Errorf("Expected: 3, got: %d", result)
	}
}

func Test_LengthOfLongestSubstring9(t *testing.T) {
	s := "abba"
	result := lengthOfLongestSubstring(s)
	if result != 2 {
		t.Errorf("Expected: 2, got: %d", result)
	}
}

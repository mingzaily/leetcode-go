package leetcode_top100

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	temp := map[int]int{}
	for i, v := range nums {
		if value, ok := temp[target-v]; ok {
			return []int{value, i}
		} else {
			temp[v] = i
		}
	}
	return nil
}

func Test_twoSum(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	for _, v := range result {
		t.Log(v)
	}
}

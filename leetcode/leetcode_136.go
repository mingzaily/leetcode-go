package main

import (
	"fmt"
	"sort"
)

// 解析，异或运算，相同值为0
func singleNumber(nums []int) int {
	value := 0
	for _, v := range nums {
		value ^= v
	}
	return value
}

// 先排序，判断挨着的值
func oldSingleNumber(nums []int) int {
	length := len(nums)

	if length <= 1 {
		return nums[0]
	}

	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if i == 0 && nums[i] != nums[i+1] {
			return nums[i]
		}
		if i == length-1 && nums[i-1] != nums[i] {
			return nums[i]
		}
		if nums[i] != nums[i+1] && nums[i-1] != nums[i] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{1}))
}

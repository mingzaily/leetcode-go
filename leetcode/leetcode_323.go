package main

import (
	"fmt"
	"sort"
)

// 哈希表 第一思路
func majorityElement(nums []int) int {
	halfLength := len(nums) / 2
	temp := make(map[int]int, halfLength)
	for _, value := range nums {
		temp[value]++
		if temp[value] > halfLength {
			return value
		}
	}
	return 0
}

// 排序
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 摩尔投票法（看解析）
func majorityElementMajor(nums []int) int {
	major, count := nums[0], 1

	for _, value := range nums[1:] {
		if count == 0 {
			count = 1
			major = value
		} else if major == value {
			count++
		} else if major != value {
			count--
		}
	}

	return major
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElementSort([]int{3, 2, 3}))
	fmt.Println(majorityElementSort([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElementMajor([]int{3, 2, 3}))
	fmt.Println(majorityElementMajor([]int{3, 3, 4}))
	fmt.Println(majorityElementMajor([]int{6, 5, 5}))
	fmt.Println(majorityElementMajor([]int{2, 2, 1, 1, 1, 2, 2}))
}

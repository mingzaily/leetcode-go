package main

import (
	"fmt"
)

// 不适用额外空间
// 将存在的值作为索引下标，将原数组对应下标的值置为负值
// 比如4存在，则将数组nums[3]，将原来值置为相反数，<0则表示存在
func findDisappearedNumbers(nums []int) []int {
	var arr []int

	for _, v := range nums {
		index := abs(v) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	for i, v := range nums {
		if v > 0 {
			arr = append(arr, i+1)
		}
	}

	return arr
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 使用额外空间
func oldFindDisappearedNumbers(nums []int) []int {
	length := len(nums) + 1
	temp := make(map[int]int, length)
	var arr []int

	for i := 1; i < length; i++ {
		temp[i] = 0
	}
	for _, v := range nums {
		temp[v]++
	}
	for key := range temp {
		if temp[key] == 0 {
			arr = append(arr, key)
		}
	}

	return arr
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}

package main

import (
	"fmt"
)

// 二分查找的前提是有序数组
// 返回下标
// 时间复杂度O(log2n)
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			right = mid - 1
		} else if target > arr[mid] {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(test, 10))
}

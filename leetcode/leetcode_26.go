package main

import "fmt"

// 快慢指针
func removeDuplicates(nums []int) int {
	fast, slow, num := 1, 1, 1
	length := len(nums)
	for ; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			// 快指针值直接覆盖慢指针指向位置，丢失数据
			nums[slow] = nums[fast]
			num++
			slow++
		}
	}
	return num
}

func main() {
	test := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(test), test)
}

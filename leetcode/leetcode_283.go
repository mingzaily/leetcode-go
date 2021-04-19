package main

import "fmt"

// 快慢指针
// 与第二十七题类似
// 当慢指针处于0元素的位置上，快指针处于非0元素的位置上，交换
func moveZeroes(nums []int) {
	fast, slow, length := 0, 0, len(nums)
	for ; fast < length; fast++ {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}

func main() {
	test := []int{0, 1, 0, 3, 12}
	moveZeroes(test)
	fmt.Println(test)
}

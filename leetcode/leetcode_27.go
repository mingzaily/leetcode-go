package main

import "fmt"

// 快慢指针
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	length := len(nums)
	for ; fast < length; fast++ {
		if nums[fast] != val {
			// 快指针值直接覆盖慢指针指向位置，丢失数据
			nums[slow] = nums[fast]
			// 快慢指针交换位置，不会丢失数据数据
			// nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fmt.Println(nums)
	}

	return slow
}

func main() {
	test := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(test, 2))
	fmt.Println(test)
}

package main

import "fmt"

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

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

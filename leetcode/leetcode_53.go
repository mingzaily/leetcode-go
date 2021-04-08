package main

import (
	"fmt"
)

// 动态规划
// f(i)代表以第i个数结尾的「连续子数组的最大和」
// f(i)=max{f(i-1)+nums[i], nums[i]}
func maxSubArray(nums []int) int {
	// dp 就是 f(i)
	dp := nums[0]
	if len(nums) <= 1 {
		return dp
	}

	max := dp
	for _, value := range nums[1:] {
		if dp+value > value {
			dp = dp + value
		} else {
			dp = value
		}
		// 保留最大值，因为f(i)不一定就是最大值
		if dp > max {
			max = dp
		}
	}

	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{0}))
}

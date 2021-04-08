package main

import "fmt"

// 暴力法 超时
func vioMaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := 0
	for i1, v1 := range prices {
		for i2, v2 := range prices {
			if i2 > i1 && v2-v1 > max {
				max = v2 - v1
			}
		}
	}

	return max
}

// 记录最低值，最大利润=最大值-最低值 即可
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minValue, maxValue := prices[0], 0
	for _, value := range prices {
		if value < minValue {
			minValue = value
		} else if value-minValue > maxValue {
			maxValue = value - minValue
		}
	}

	return maxValue
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 4, 2}))
	fmt.Println(maxProfit([]int{2, 4, 1}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3}))
}

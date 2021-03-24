package main

import (
	"fmt"
)

// 斐波那契
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return climbStairs(n-1) + climbStairs(n-2)
	}
}

// 动态规划1 存储所有算过
func climbStairs2(n int) int {
	dep := []int{1, 1}
	for i := 2; i <= n; i++ {
		dep = append(dep, dep[i-1]+dep[i-2])
	}
	return dep[n]
}

// 动态规划2 临时变量来代替f(n-1)和f(n-2)
func climbStairs3(n int) int {
	// 0->0 1步
	// 0->1 1步
	prev, cur := 1, 1
	for i := 2; i <= n; i++ {
		now := cur + prev
		prev = cur
		cur = now
	}
	return cur
}

func main() {
	//fmt.Println(climbStairs(50)) 太耗时了
	fmt.Println(climbStairs2(50))
	fmt.Println(climbStairs3(50))
}

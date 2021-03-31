package main

import "fmt"

// 当前台阶消耗的体力一定取决于前面所消耗的体力，可以选择从第0个台阶或者第1个台阶走，可以走一步或者两步
// 状态转移方程 f(x) = min(f(x-1) + cost[x], f(x-2) + cost[x])
// 解释 当前台阶消耗体力取 我上两步+到这个台阶消耗的量或我上一步+到这个台阶消耗的量 其中的最小值
//              终点
//              ____
//         20  |
//          ___|   第2个台阶
//     15  |
//      ___|  第1个台阶
// 10  |
// ____|  第0个台阶
// 终点消耗量其实就是比较 从第1个台阶上来消耗的消耗量最小还是从第2个台阶上来消耗量最小
// min(dp[2], dp[1]) = min(30, 15) = 15     不在循环里
// 第2个台阶消耗量取决于从 第1个台阶上来消耗的小还是从0个台阶上来消耗的量小
// dp[2] = min(dp[1], dp[0]) + 20  = min(15, 10) + 20= 30
// 第1个台阶消耗量取决于从 第0个台阶上来消耗的小还是直接踏上来小(边界值)
// dp[1] = min(dp[0], dp[-1]) + 15 = min(10, 0) + 15 = 15 = cost[1]
// 第0个台阶消耗量(边界值)
// dp[0] = min(dp[-1, dp[-2]]) + 10 = min(0, 0) + 10 = 10 = cost[0]

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = 0, cost[1]
	for i := 2; i < len(cost); i++ {
		if dp[i-1] < dp[i-2] {
			dp[i] = dp[i-1] + cost[i]
		} else {
			dp[i] = dp[i-2] + cost[i]
		}
	}
	fmt.Println(dp)
	if dp[len(cost)-1] < dp[len(cost)-2] {
		return dp[len(cost)-1]
	}
	return dp[len(cost)-2]
}

// 空间优化
func minCostClimbingStairs2(cost []int) int {
	lastTwo, lastOne := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		if lastOne < lastTwo {
			lastTwo, lastOne = lastOne, lastOne+cost[i]
		} else {
			lastTwo, lastOne = lastOne, lastTwo+cost[i]
		}
	}
	if lastOne < lastTwo {
		return lastOne
	}
	return lastTwo
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15}))
}

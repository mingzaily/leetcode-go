/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	maxLen := 1
	currentLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			currentLen++
		} else if nums[i] != nums[i-1] {
			currentLen = 1
		}
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

// @lc code=end


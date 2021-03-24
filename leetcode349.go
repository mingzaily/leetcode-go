package main

import (
	"fmt"
	"sort"
)

func intersection1(nums1 []int, nums2 []int) (arr []int) {
	temp := map[int]int{}
	for _, v := range nums1 {
		temp[v] = 0
	}
	for _, v := range nums2 {
		if value, ok := temp[v]; ok && value == 0 {
			arr = append(arr, v)
			temp[v]++
		}
	}
	return arr
}

// 转换成有序切片
// 双指针
func intersection2(nums1 []int, nums2 []int) (arr []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for index1, index2 := 0, 0; index1 < len(nums1) && index2 < len(nums2); {
		if nums1[index1] == nums2[index2] {
			// 不重复，缓存的最后一位不能大于等于当前值
			if arr == nil || arr[len(arr)-1] < nums1[index1] {
				arr = append(arr, nums1[index1])
			}
			index1++
			index2++
		} else if nums1[index1] < nums2[index2] {
			index1++
		} else {
			index2++
		}
	}
	return arr
}

func main() {
	fmt.Println(intersection1([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersection2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

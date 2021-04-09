package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	point1, point2 := 0, 0
	len1, len2 := len(nums1), len(nums2)
	temp := make([]int, len1+len2)
	for i := range temp {
		if point1 < len1 && (point2 == len2 || nums1[point1] < nums2[point2]) {
			temp[i] = nums1[point1]
			point1++
		} else {
			temp[i] = nums2[point2]
			point2++
		}
	}

	if len(temp)%2 == 0 {
		return float64(temp[len(temp)/2-1]+temp[len(temp)/2]) / 2
	} else {
		return float64(temp[len(temp)/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

package main

import "fmt"

// 基于快速排序思路的快速选择
func getLeastNumbers(arr []int, k int) []int {
	length := len(arr)
	if k > length {
		return []int{}
	}
	quickSelect(arr, 0, length-1, k)
	return arr[:k]
}

func quickSelect(arr []int, startIndex, endIndex, k int) {
	if startIndex < endIndex {
		// 进行切分
		pivotIndex := partition(arr, startIndex, endIndex)
		if pivotIndex == k {
			return
		} else if k < pivotIndex {
			quickSelect(arr, startIndex, pivotIndex-1, k)
		} else {
			quickSelect(arr, pivotIndex+1, endIndex, k)
		}
	}
}

func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex

	for left != right {
		for left < right && pivot < arr[right] {
			right--
		}
		for left < right && pivot >= arr[left] {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[startIndex], arr[left] = arr[left], arr[startIndex]

	return left
}

func main() {
	fmt.Println(getLeastNumbers([]int{0, 7, 2, 4, 3}, 4))
}

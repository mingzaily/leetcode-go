package main

import "fmt"

func quickSort(arr []int, startIndex, endIndex int) {
	if startIndex < endIndex {
		// 进行切分
		pivotIndex := partition(arr, startIndex, endIndex)
		// 左数组
		quickSort(arr, startIndex, pivotIndex-1)
		// 右数组
		quickSort(arr, pivotIndex+1, endIndex)
	}
}

func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex

	for left != right {
		// 大于基准数，右指针左移
		for left < right && pivot < arr[right] {
			right--
		}
		// 小于基准数，左指针右移
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
	test := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	quickSort(test, 0, len(test)-1)
	fmt.Println(test)
}

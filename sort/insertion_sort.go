package main

import "fmt"

// 直接插入排序
func directInsertionSort(arr []int) {
	for i := range arr {
		preIndex := i
		currentValue := arr[preIndex]
		for preIndex > 0 && arr[preIndex-1] > currentValue {
			arr[preIndex] = arr[preIndex-1]
			preIndex--
		}
		arr[preIndex] = currentValue
	}
}

func main() {
	test := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	directInsertionSort(test)
	fmt.Println(test)
}

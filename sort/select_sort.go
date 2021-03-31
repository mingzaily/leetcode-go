package main

import "fmt"

func selectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	test := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	selectSort(test)
	fmt.Println(test)
}

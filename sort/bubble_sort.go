package main

import "fmt"

func bubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	test := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	bubbleSort(test)
	fmt.Println(test)
}

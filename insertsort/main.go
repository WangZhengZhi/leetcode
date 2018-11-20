package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 99, 1, 45, 87, 111}
	InsertSort(arr)
	fmt.Println(arr)
}
func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]

			}
		}
	}
}

package main

import (
	"fmt"
)

/*非递归版本的二分查找*/

func BinarySearch(arr []int, low, high, target int) int {
	mid := low + (high-low)/2
	for low <= high {
		if arr[mid] == target {
			return mid

		} else if arr[mid] > target {
			high = mid

		} else if arr[mid] < target {
			low = mid

		}
	}

	return -1
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(BinarySearch(arr,0,7,4))
}

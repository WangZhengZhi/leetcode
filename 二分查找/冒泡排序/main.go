package main

import "fmt"

func main() {
	arr := []int{1, 23, 0, 3, 67, 88}
	Bsort(arr)
	fmt.Println(arr)
}
func Bsort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] //升序排列
				//降序： if arr[j]<arr[j+1]
			}

		}

	}
}

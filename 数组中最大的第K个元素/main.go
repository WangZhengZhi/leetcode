package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(a, 2))

}
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	target := nums[len(nums)-k]
	return target
}

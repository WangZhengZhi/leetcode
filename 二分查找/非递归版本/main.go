package main

/*非递归版本*/
func main() {

}
func Binarysearch(arr []int, high, low, target int) int {
	mid := low + (high-low)/2
	if arr[mid] == target {
		return mid

	} else if arr[mid] > target {
		high = mid
	} else if arr[mid] < target {
		low = mid

	}
	return 0

}

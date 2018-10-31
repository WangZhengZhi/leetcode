package main

/*非递归版本的二分查找*/
func main() {

}
func BinarySearch(arr []int, low, high, target int) int {
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

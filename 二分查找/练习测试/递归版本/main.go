package main

/*binary search 递归版本*/
func main() {

}
func BinarySearch(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if arr[mid] > target {
		return BinarySearch(arr, low, mid-1, target)

	}
	if arr[mid] < target {
		return BinarySearch(arr, mid+1, high, target)

	}
	return mid
}

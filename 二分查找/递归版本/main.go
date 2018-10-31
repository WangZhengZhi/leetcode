package main

/*递归版本*/
func main() {

}
func Binarysearch(arr []int, low, high, target int) int {

	if low > high {
		return -1

	}
	mid := low + (high-low)/2
	if arr[mid] > target {
		return Binarysearch(arr, low, mid-1, target)
	} else if arr[mid] < target {
		return Binarysearch(arr, mid+1, high, target)

	}
	return mid
}

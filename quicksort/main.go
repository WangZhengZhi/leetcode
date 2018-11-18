package main

import "fmt"

func main() {
	arr := []int{5, 9, 1, 2, 0, 88}
	QuickSort(arr)
	fmt.Println(arr)
}
func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return

	}
	mid, i := arr[0], 1       //分水岭，右侧元素下标为1
	head, tail := 0, length-1 //头/尾部
	for head < tail {
		if arr[i] > mid { // 如果分水岭右侧的元素大于分水岭，就将右侧的尾部元素和分水岭右侧元素交换
			arr[i], arr[tail] = arr[tail], arr[i]
			tail-- //尾部下标左移动

		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++ //头部右移动
			i++    //下标右移动
		}
	}
	//分水岭左右做递归处理
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])

}

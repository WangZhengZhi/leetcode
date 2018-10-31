package main

import (
	"fmt"
)

/*冒泡排序*/
func main() {
	a := []int{1, 6, 3, 5, 7}
	Bsort(a)
	Selectsort(a)

}
func Bsort(a []int) {

	for i := 0; i < len(a); i++ {

		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]

			}
		}
	}
	fmt.Println(a)

} /*冒泡*/
func Quicksort(a []int) {

}
func Selectsort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				a[j], a[min] = a[min], a[j]

			}

		}

	}
	fmt.Println(a)

} //选择排序

package main

import (
	"fmt"
	"math"
)

// 汉诺塔问题
func main() {

}
func Hannuoi(n int, x, y, z string) int {
	if n == 0 {
		fmt.Println("不需要移动")
		return 0
	}
	if n == 1 {
		fmt.Println(n, "move ", x, "to ", y)
	} else {
		Hannuoi(n-1, x, y, z)
		fmt.Println(n, "move ", x, "to ", y)
		Hannuoi(n-1, y, x, z)

	}
	return int(math.Pow(2, float64(n)) - 1)
}

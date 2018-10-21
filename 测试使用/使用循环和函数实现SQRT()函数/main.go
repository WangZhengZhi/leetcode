package main

import (
	"fmt"
	"math"
)

/*package 使用循环和函数实现SQRT__函数
z=z-(z^2-x)/2z
多次使用可以逼近函数
牛顿法逼近sqrt()函数
*/
func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}
func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 1000; i++ {
		z = z - (z*z-x)/(2*z)

	}
	return z
}

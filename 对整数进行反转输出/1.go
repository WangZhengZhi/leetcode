package main

import (
	"fmt"
	"math"
)

/*对整数进行反转输出*/
/*x%10取出X的最后一位*/
func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if (num > math.MaxInt32) || (num < (-math.MaxInt32)) {
		return 0
	}
	return
}
func main() {
	var x int
	x = -32
	fmt.Println(reverse(x))
}

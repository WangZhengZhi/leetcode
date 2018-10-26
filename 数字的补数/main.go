package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 1
	fmt.Println(findComplement(num))
}
func findComplement(num int) int {
	divnum := -1
	modnum := -1
	sum := 0
	var n int = 0
	for divnum != 0 {
		divnum = num / 2
		modnum = num % 2
		num = divnum
		sum += (1 - modnum) * int(math.Pow(float64(2), float64(n)))
		n++
	}
	return sum

}

package main

import (
	"fmt"
	"math"
)

/*if条件语句问题*/
func main() {
	value := Pow(1, 2, 5)
	fmt.Println(value)
}
func Pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		return limit
	}
}

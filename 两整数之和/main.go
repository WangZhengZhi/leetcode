package main

import "fmt"

func main() {
	fmt.Println(getSum(-1, 2))
	fmt.Println((-1)^4)

}
func getSum(a int, b int) int {

	result := a ^ b
	carry := (a & b) << 1
	for carry != 0 {
		return getSum(result, carry)

	}
	return result

}

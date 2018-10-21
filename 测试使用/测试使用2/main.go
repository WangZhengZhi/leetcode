package main

import "fmt"

/*关于移位运算问题*/
func main() {
	var i int
	i = 0x80
	j := i << 1
	fmt.Println(j)
	fmt.Println(i)
}

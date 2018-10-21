package main

import (
	"fmt"
	"math"
)

func main() {
	var i bool
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	var num int
	var f float64 = float64(num)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", num)
	const num1 = 11.0
	fmt.Printf("%T\n", num1)
	const num2 = 10 << 2 //10=1010-----左移动2位--  101000=40
	fmt.Printf("%d\n", num2)
	fmt.Printf("%T\n", num2)   //打印数据类型
	fmt.Println(math.MaxInt32) //最大的32位int数字
	fmt.Println(math.MaxInt64) //最大的64位int数字
}

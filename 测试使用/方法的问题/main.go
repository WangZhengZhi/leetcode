package main

import "fmt"

/*go没有类但是有方法*/
type Myfloat float64

func (f Myfloat) Abs() float64 { //func(形参 形参类型)函数名（）函数类型{}
	if f < 0 {
		return float64(-f)

	}
	return float64(f)
}
func main() {
	var num float64 = -1.3
	fmt.Println(Myfloat(num).Abs()) //注意如果形参是指针类型，记得加& 取地址符号
}

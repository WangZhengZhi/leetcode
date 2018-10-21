package main

import (
	"fmt"
	"strconv"
)

/*结构体与方法练习*/
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func BreakErr(e error) {
	if e != nil {
		fmt.Println("error=", e)
	}
} //处理错误
func main() {
	a := Person{"Joel", 31}  //实现了person中的string（）方法
	z := Person{"Smith", 45} //实现了person中的string（）方法
	fmt.Println(a, z)
	num, err := strconv.Atoi("12345")
	BreakErr(err)
	fmt.Printf("%T\n", num)

	num1 := strconv.Itoa(123456789)
	fmt.Printf("num1 type:%T\n", num1)
	fmt.Printf("num1:=%s\n", num1)

}

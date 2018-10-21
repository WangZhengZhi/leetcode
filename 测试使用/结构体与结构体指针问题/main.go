package main

import "fmt"

/*结构体与结构体指针问题*/
type student struct {
	Id   int
	Name string
}

func main() {
	su := student{21, "zhenhzhiwang"}
	fmt.Println(su.Id)
	fmt.Println(su.Name)
	p := &su //指针是地址操作，引用传递，不是值传递
	p.Name = "hello"
	p.Id = 22
	fmt.Println(*p) //打印的结果不带&
	fmt.Println(p)  //打印的结果带&
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 1)
	slice = append(slice, 1)
	slice = append(slice[:1], slice[2:]...)
	fmt.Println(slice)
	m := map[int]string{1: "hello", 2: "world"}
	for key, value := range m {
		if key == 2 {
			fmt.Print("*--*")
		}
		fmt.Println(key, value)
	}
	m[3] = "你好"
	fmt.Println(m[3])
	key, ok := m[1]
	fmt.Println(key, ok)
	delete(m, 1)
}

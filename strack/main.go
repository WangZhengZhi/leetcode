package main

import (
	"fmt"
)

type Element interface {
}
type Stack struct {
	Data []Element
}

//生成堆栈
func Makestack() *Stack {
	return &Stack{}
}

//判断栈的容量
func (s *Stack) Size() int {
	return len(s.Data)
}

//判断是否为空栈
func (s *Stack) Isempty() bool {
	if s.Size() == 0 {
		return true

	} else {
		return false
	}

}

//push 入栈
func (s *Stack) Push(data Element) {
	s.Data = append(s.Data, data)
}

//pop出栈
func (s *Stack) Pop() Element {
	if s.Size() == 0 {
		fmt.Println("empty stack！！！！")
		return nil

	} else {
		lastdata := s.Data[s.Size()-1] //取出最后一个数据
		s.Data[s.Size()-1] = nil       //将最后一个元素设置为nil
		s.Data = s.Data[:s.Size()-1]   //截取除了最后一个元素的所有元素
		return lastdata                //返回最后一个元素,先进后出

	}
}

//弹出最顶端的数据,但是不删除元素
func (s *Stack) Top() Element {
	if s.Size() == 0 {
		fmt.Println("is empty!!!")
		return nil

	} else {
		return s.Data[s.Size()-1]
	}
}

func main() {
	stack := Makestack()
	for i := 0; i < 9; i++ {
		stack.Push(i)
	}
	fmt.Println("starck:", stack)
	/*fmt.Println("top:", stack.Top())*/
	fmt.Println(stack.Pop())
	fmt.Println("starck:", stack)


}

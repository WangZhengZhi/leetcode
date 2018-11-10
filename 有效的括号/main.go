package main

import "fmt"

func main() {
	var s string
	s = "]"
	fmt.Println(isValid(s))

}

type Stack struct {
	Data []rune
}

func MakeEmpty() *Stack {
	return &Stack{}
}
func (s *Stack) Size() int {
	return len(s.Data)
}
func (s *Stack) Pop() rune {
	if s.Size() == 0 {
		return 0
	} //特殊情况
	lastdata := s.Data[s.Size()-1]
	s.Data = s.Data[:s.Size()-1]
	return lastdata
}
func (s *Stack) Push(data rune) {
	s.Data = append(s.Data, data)
}
func (s *Stack) Isempty() bool {
	return s.Size() == 0
}
func isValid(s string) bool {

	stack := MakeEmpty()
	runes := []rune(s)
	for _, value := range runes {

		switch value {
		case '[':
			stack.Push(']')

		case '{':
			stack.Push('}')

		case '(':
			stack.Push(')')

		case '}', ']', ')':
			if stack.Pop() != value { //如果出栈的值和入栈的值不一样，则为错误
				return false

			}

		}
	}
	if stack.Size() > 0 { //出栈完成后，理论上正确的输入长度等于0
		return false

	}
	return true

}

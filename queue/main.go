package main

import (
	"fmt"
)

type Element interface {
}

type Queue struct {
	Data []Element
}

//生成空队列
func Makequeue() *Queue {
	return &Queue{}
}

//判断队列的size
func (q *Queue) Size() int {
	return len(q.Data)
}

//判断队列是否为空
func (q *Queue) Isempty() bool {
	return len(q.Data) == 0
}

//入队列 push
func (q *Queue) Push(data Element) {

	q.Data = append(q.Data, data)
}

//出队列,并返回出队列的数据
func (q *Queue) Pop() Element {
	if q.Size() == 0 {
		return nil

	}
	firstdata := q.Data[0]
	q.Data = q.Data[1:q.Size()]
	return firstdata
}
func main() {
	queue := Makequeue()

	for i := 0; i < 10; i++ {
		queue.Push(i)

	}
	fmt.Println(queue)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Isempty())
	fmt.Println(queue.Size())
}

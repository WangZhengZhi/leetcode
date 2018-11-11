package main

import (
	"testing"
)

var queue = Makequeue()

func TestMakequeue(t *testing.T) {

	if queue.Data == nil {
		t.Error("测试通过")

	} else {
		t.Error("不通过")
	}
}
func TestQueue_Push(t *testing.T) {
	for i := 0; i < 100; i++ {
		queue.Push(i)
	}
	for i := 0; i < 100; i++ {
		if queue.Pop() != i {
			t.Error("测试不通过","错误的数据是：",i)

		} else {
			t.Error(i," 测试通过 ")
		}

	}

}
func TestQueue_Size(t *testing.T) {
	for i := 0; i < 100; i++ {
		queue.Push(i)

	}
	if queue.Size()==len(queue.Data) {
		t.Error("测试通过")

	}else{
		t.Error("测试不通过")
	}
}

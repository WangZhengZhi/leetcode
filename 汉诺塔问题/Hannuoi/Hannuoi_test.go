package main

import "testing"

func TestHannuoi(t *testing.T) {
	num := Hannuoi(5, "x", "y", "z")
	if num != 31 {
		t.Error("测试失败")
	} else {
		t.Error("测试成功")
	}

}

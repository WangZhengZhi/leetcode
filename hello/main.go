package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	_"hello/models"/*调用models中的init函数*/
)

func main() {
	beego.Run()
}


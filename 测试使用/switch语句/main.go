package main

import (
	"fmt"
	"runtime"
	"time"
)

/*switch语句的使用问题*/
func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os is macos")
	case "freedsb":
		fmt.Println("os is freedsb")
	case "linux":
		fmt.Println("os is linux")
	case "windows":
		fmt.Println("os is windos")
		fallthrough //fallthrough有穿透的效果，可以执行到下面一个分支

	case "hehe":
		fmt.Println("hehe")
	default:
		fmt.Println("os is ", os)

	}
	/*switch语句没有条件表达式的情况下，相当于switch true{}每一个分支的条件都是BOOL值，执行true的分支
	或者执行default*/
	switch { //switch语句之后是没有表达式的
	case time.Now().Hour() > 8 && time.Now().Hour() < 12:
		fmt.Println("早上好")
	case time.Now().Hour() > 12:
		fmt.Println("中午好")
	case time.Now().Hour() > 17:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")
	}
}

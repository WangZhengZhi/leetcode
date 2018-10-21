package main

import (
	"fmt"
	"io"
	"strings"
)

/*readers读取数据流的问题*/
func CheckErr(e error) {
	if e != nil {
		fmt.Println("error:", e)
	}

} //检查并打印error
func main() {
	r := strings.NewReader("Hello,Reader!")
	b := make([]byte, 8) //每次读取8个字节
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n]) //字节切片使用%q
		if err == io.EOF {              //读取到最后一个字节的时候
			break
		}
	}

}

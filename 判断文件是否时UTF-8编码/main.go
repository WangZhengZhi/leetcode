package main

import (
	"fmt"
	"io/ioutil"
)

/*判断文件是不是UTF8文件*/
/*unicode编码原则问题
utf8是一个可变的长的编码形式
字符长度范围是1-4（包含4）
一字节：0*******

两字节：110*****，10******

三字节：1110****，10******，10******

四字节：11110***，10******，10******，10******



*/
func main() {

	buf, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("是否是UTF8编码：", IsUtf8(buf))
	}

}
func IsUtf8(buf []byte) bool {
	switch len(buf) {
	case 1:
		if buf[0]&0x80 == 0 {
			return true
		} else {
			return false
		}
	case 2:
		if buf[0]<<2&0x80 == 0 && buf[1]<<1&0x80 == 0 {
			return true
		} else {
			return false
		}
	case 3:
		if buf[0]<<3&0x80 == 0 && buf[1]<<1&0x80 == 0 && buf[2]<<1&0x80 == 0 {
			return true
		} else {
			return false
		}
	case 4:
		if buf[0]<<4&0x80 == 0 && buf[1]<<1&0x80 == 0 && buf[2]<<1&0x80 == 0 && buf[3]<<1&0x80 == 0 {
			return true
		} else {
			return false
		}
	default:
		return false

	}
}

package main

import "fmt"

/*题目：回文数*/
/*判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？*/
//说明：此题目1位数默认是回文数字，并且负数不是回文数字
func isPalindrome(x int) (flag bool) {
	tmp := x //临时保存X的值。便于以后比较
	var num int = 0
	for x != 0 && tmp > 0 { //过滤负数
		num = num*10 + x%10
		x = x / 10
	}
	return tmp == num
}
func main() {
	var x int = 1221
	fmt.Println(isPalindrome(x))

}

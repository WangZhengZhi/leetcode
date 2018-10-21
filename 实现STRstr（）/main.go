package main

import "fmt"

/*给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
*/
func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))

}
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ { //快指针
		for j := 0; ; j++ { //慢指针
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break

			}
		}
	}
}

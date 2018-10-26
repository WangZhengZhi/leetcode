package main

import "fmt"

func main() {
s:="hello,nihao"
fmt.Println(reverseString(s))
}
func reverseString(s string) string {
	runs := []rune(s)
	for i,j:=0,len(runs)-1;i<j;i,j=i+1,j-1{
		runs[i],runs[j]=runs[j],runs[i]

	}
	return string(runs)
}

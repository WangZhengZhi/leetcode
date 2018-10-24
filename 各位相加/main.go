package main

import "fmt"

/*各位相加*/
func main(){
fmt.Println(addDigits(19))
}
func addDigits(num int) int {
var sum int =0
	for  {
		sum+=num%10
		num=num/10
		if num==0{
			break
		}
/*此时的sum为各个位数之和*/

	}
if sum>=10{/*如果大于等于10执行递归函数继续计算*/
	num=sum
	return addDigits(num)
}

return sum
}

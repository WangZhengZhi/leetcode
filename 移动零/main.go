package main

import (
	"fmt"
)

func main() {

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
func moveZeroes(nums []int) {
	/*if nums==nil||len(nums)==0{
		return
	}
	i:=0
	for _,value:=range nums{
		if value!=0{
			nums[i]=value
			i++
		}
	}
	for i<len(nums)  {
		nums[i]=0
		i++

	}*/
	var i int=0
	var j int=0
	for j<len(nums)  {
		if nums[j]!=0 {
			nums[i],nums[j]=nums[j],nums[i]//nums[j]的值不为0，则向前移动
			i++
		}
		j++

	}

}

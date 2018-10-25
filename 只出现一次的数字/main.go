package main

import "fmt"

/*只出现一次的数字*/
func main() {
	nums := []int{1, 4, 4}
	fmt.Println(singleNumber(nums))
	fmt.Println(1 ^ 4 ^ 1)
}
func singleNumber(nums []int) int {
	var ret int = 0
	for _, value := range nums {
		ret = ret ^ value
		/*
		^ 为异或运算，相同为0，
			a ^ a = 0
			a ^ 0 = a
			a ^ b ^ c = a ^ (b ^ c)
			a^b^a=b
		*/
	}
	return ret

}
/*func singleNumber(nums []int) int {
    count:=0
    var re int
    for i:=0;i<len(nums);i++{
        count = 0
        for j:=0;j<len(nums);j++{
            if nums[i] == nums[j] && i!=j{
                count++
            }
        }
        if count == 0{
            re= nums[i]
        }

    }
    return re
}*/
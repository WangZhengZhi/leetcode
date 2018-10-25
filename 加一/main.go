package main

import "fmt"

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
}
func plusOne(digits []int) []int {
	nums := make([]int, len(digits))
	copy(nums, digits)
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		num := (nums[i] + carry) % 10  //进位，满10则为0，否则还是自己
		carry = (nums[i] + carry) / 10 //判断是不是超出10，如果超出，则执行if语句将carry的值赋值到一位

		nums[i] = num

		// if there is no more carry, we can terminate the loop early
		/*if carry == 0 {
			return nums
		}*/
	}

	if carry != 0 {
		nums = append([]int{carry}, nums...)
	}

	return nums
}

package main

import "fmt"

/*给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用
数组中的元素是可以重复出现的
*/
func twoSum(nums []int, target int) (newnum []int) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i && j > 0; j-- {
			if nums[i]+nums[j] == target {
				newnum = append(newnum, i)
				newnum = append(newnum, j)
				break
			}
		}

	}

	return

}
func main() {
	nums := []int{3, 2, 3}
	for _, value := range twoSum(nums, 6) {
		fmt.Println(value)
	}
}

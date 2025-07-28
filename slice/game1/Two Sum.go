package main

import "fmt"

/*
** 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
** 输入：nums = [2,7,11,15], target = 9
 */

func main() {
	testslice := []int{2, 7, 11, 15}
	target := 9
	var target2 int = target
	fmt.Println(twoSum(testslice, target2))
}
func twoSum(nums []int, target int) []int {
	//result := make([]int,2)
	result := make([]int, 0, 2)
	for index1, slow := range nums {
		for index2, fast := range nums {
			if fast+slow == target {
				toadd := []int{index1, index2}
				result = append(result, toadd...)
				fmt.Println(result)
				break
			}
		}
	}
	return result
}

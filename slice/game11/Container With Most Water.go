package main

import "fmt"

/* 7.28
** 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
** Input: [1,8,6,2,5,4,8,3,7]
** Output: 49
 */

func main() {
	testslice := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	target := 49
	fmt.Println(mostWater_on2(testslice), mostWater_on(testslice))
	fmt.Println(mostWater_on(testslice) == target)
}

func mostWater_on2(nums []int) int {
	result := 0
	for slow, num1 := range nums {
		for fast := len(nums) - 1; fast < len(nums); fast++ {
			num2 := nums[fast]
			result = max(result, (fast-slow)*min(num1, num2))
		}
	}

	return result
}
func mostWater_on(nums []int) int {
	result, slow, fast := 0, 0, len(nums)-1
	for slow < fast {
		high := min(nums[slow], nums[fast])
		width := fast - slow
		if nums[slow] < nums[fast] {
			high = nums[slow]
			slow++
		} else {
			high = nums[fast]
			fast--
		}
		result = max(result, width*high)
	}

	return result
}

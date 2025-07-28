package game11

import "fmt"

/* 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
** Input: [1,8,6,2,5,4,8,3,7]
**Output: 49
 */

func main() {
	testslice := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	target := 49
	fmt.Println(twoSum(testslice, target))
}

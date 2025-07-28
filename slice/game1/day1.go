package main

import "fmt"

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

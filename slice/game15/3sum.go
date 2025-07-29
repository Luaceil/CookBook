package main

import "fmt"

/*
	给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
	Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[

	[-1, 0, 1],
	[-1, -1, 2]

]
*/
type Set[T comparable] struct {
	data map[T]struct{} // 用 struct{} 作为值，不占用内存
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		data: make(map[T]struct{}),
	}
}
func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{} // 用空结构体占位 struct{}空类型 {}初始化为空
}
func (s *Set[T]) Contains(v T) bool {
	_, exists := s.data[v]
	return exists
}

func main() {
	testslice := []int{-1, 0, 1, 2, -1, -4}
	//target := 49
	fmt.Println(sumof3(testslice))
	//fmt.Println(sumof3(testslice) == target)
}

func sumof3(nums []int) [][]int {
	result := make([][]int, 0)
	object := make(map[int][]int)
	for slow := 0; slow < len(nums); slow++ {
		for fast := slow + 1; fast < len(nums); fast++ {
			object[nums[slow]+nums[fast]] = append(object[nums[slow]+nums[fast]], slow, fast)
		}
	}
	for slow, num := range nums {
		if object[-num] != nil {
			for index := 0; index < len(object[-num]); index += 2 {
				if slow < object[-num][index] {
					result = append(result, []int{slow, object[-num][index], object[-num][index+1]})
				}
			}
		}
	}
	return result
}

// +build exercise17

package main

import "fmt"

func farPair(nums []int) []int {
	var result []int
	max := nums[0]
	min := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return append(result, max, min)
}

func main() {
	var nums []int
	nums = append(nums, 2, 7, 11, 15)
	fmt.Println(farPair(nums))
}

// +build twoSumFaster

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	exist := make(map[int]int)
	for i, v := range nums {
		if exist[target-v] != 0 {
			result = append(result, exist[target-v]-1, i)
			break
		} else {
			exist[v] = i + 1
		}
	}
	return result
}

func main() {
	var nums []int
	nums = append(nums, 2, 7, 11, 15)
	target := 9
	fmt.Println(twoSum(nums, target))
}

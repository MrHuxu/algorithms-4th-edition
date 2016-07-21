// +build selectSort

package main

import "fmt"

func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		tmp := nums[i]
		nums[i] = nums[min]
		nums[min] = tmp
	}
}

func main() {
	var nums []int
	nums = append(nums, 2, 1, 3, 5, 2, 4, 1)
	selectSort(nums)
	fmt.Println(nums)
}

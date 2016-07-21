// +build insertSort

package main

import "fmt"

func insertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
}

func main() {
	var nums []int
	nums = append(nums, 7, 12, 2, 1, 3, 5, 2, 4, 1)
	insertSort(nums)
	fmt.Println(nums)
}

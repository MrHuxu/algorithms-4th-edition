// +build shellSort

package main

import "fmt"

func shellSort(nums []int) {
	N := len(nums)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && nums[j] < nums[j-1]; j -= h {
				tmp := nums[j]
				nums[j] = nums[j-h]
				nums[j-1] = tmp
			}
		}
		h = h / 3
	}
}

func main() {
	var nums []int
	nums = append(nums, 7, 12, 2, 1, 3, 5, 2, 4, 1)
	shellSort(nums)
	fmt.Println(nums)
}

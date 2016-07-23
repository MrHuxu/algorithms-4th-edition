// +build mergeSortBu

package main

import "fmt"

var aux []int

func merge(nums []int, lo int, mid int, hi int) {
	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		aux[k] = nums[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > hi {
			nums[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			nums[k] = aux[j]
			j++
		} else {
			nums[k] = aux[i]
			i++
		}
	}
}

func mergeSortBU(nums []int) {
	for _, v := range nums {
		aux = append(aux, v)
	}

	sortDetail(nums, 0, len(nums)-1)
}

func sortDetail(nums []int, lo int, hi int) {
	N := len(nums)
	for sz := 1; sz < N; sz *= 2 {
		for lo := 0; lo < N-sz; lo += sz * 2 {
			var tmp int
			if lo+sz*2 < N {
				tmp = lo + sz*2 - 1
			} else {
				tmp = N - 1
			}
			merge(nums, lo, lo+sz-1, tmp)
		}
	}
}

func main() {
	var nums []int
	nums = append(nums, 4, 3, 1, 3, 2, 5)
	mergeSortBU(nums)
	fmt.Println(nums)
}

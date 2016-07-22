// +build mergeSort

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

func mergeSort(nums []int) {
	for _, v := range nums {
		aux = append(aux, v)
	}

	sortDetail(nums, 0, len(nums)-1)
}

func sortDetail(nums []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := (lo + hi) / 2
	sortDetail(nums, lo, mid)
	sortDetail(nums, mid+1, hi)
	merge(nums, lo, mid, hi)
}

func main() {
	var nums []int
	nums = append(nums, 4, 3, 1, 3, 2, 5)
	mergeSort(nums)
	fmt.Println(nums)
}

// +build merge

package main

import "fmt"

var aux []int

func merge(nums []int, lo int, mid int, hi int) {
	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		aux = append(aux, nums[k])
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

func main() {
	var nums []int
	nums = append(nums, 4, 3, 2, 1, 4, 2)
	merge(nums, 0, 3, 5)
	fmt.Println(nums)
}

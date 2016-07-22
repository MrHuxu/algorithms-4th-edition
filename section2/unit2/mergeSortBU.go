// +build mergeSortBu

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

}

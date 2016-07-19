// +build exercise21

package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func find(h *listNode, k int) *listNode {
	count := 1
	for ; h.Next != nil; h = h.Next {
		if count == k {
			break
		} else {
			count++
		}
	}
	return h
}

func main() {
	f := listNode{6, nil}
	e := listNode{5, &f}
	d := listNode{4, &e}
	c := listNode{3, &d}
	b := listNode{2, &c}
	a := listNode{1, &b}
	fmt.Println(find(&a, 1))
	fmt.Println(find(&a, 3))
	fmt.Println(find(&a, 5))
}

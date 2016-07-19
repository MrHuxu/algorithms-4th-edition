// +build exercise25

package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func insertAfter(h *listNode, k int, v int) {
	count := 1
	for ; h.Next != nil; h = h.Next {
		if count == k {
			tmp := listNode{v, h.Next}
			h.Next = &tmp
			break
		} else {
			count++
		}
	}
}

func main() {
	f := listNode{6, nil}
	e := listNode{5, &f}
	d := listNode{4, &e}
	c := listNode{3, &d}
	b := listNode{2, &c}
	a := listNode{1, &b}
	insertAfter(&a, 1, 0)
	fmt.Println(a.Next.Next)
}

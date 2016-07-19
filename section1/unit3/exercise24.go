// +build exercise21

package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func removeAfter(h *listNode, k int) {
	count := 1
	for ; h.Next != nil; h = h.Next {
		if count == k {
			h.Next = h.Next.Next
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
	removeAfter(&a, 1)
	fmt.Println(a.Next)
	removeAfter(&a, 2)
	fmt.Println(a.Next.Next)
}

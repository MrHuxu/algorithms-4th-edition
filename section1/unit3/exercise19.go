// +build exercise19

package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func deleteTail(h *listNode) {
	tmp := h
	for ; h.Next != nil; h = h.Next {
		tmp = h
	}
	tmp.Next = nil
}

func main() {
	f := listNode{6, nil}
	e := listNode{5, &f}
	d := listNode{4, &e}
	c := listNode{3, &d}
	b := listNode{2, &c}
	a := listNode{1, &b}
	fmt.Println(a.Val)
	fmt.Println(a.Next)
	fmt.Println(a.Next.Next.Next.Next.Next)
	deleteTail(&a)
	fmt.Println(a)
	fmt.Println(a.Next)
	fmt.Println(a.Next.Next.Next.Next.Next)
}

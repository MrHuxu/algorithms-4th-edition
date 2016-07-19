// +build exercise26

package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func remove(h *listNode, v int) {
	tmpHead := &listNode{0, h}
	for ; tmpHead != nil; tmpHead = tmpHead.Next {
		if tmpHead.Next != nil && tmpHead.Next.Val == v {
			tmpHead.Next = tmpHead.Next.Next
		}
	}
}

func main() {
	f := listNode{2, nil}
	e := listNode{5, &f}
	d := listNode{2, &e}
	c := listNode{3, &d}
	b := listNode{2, &c}
	a := listNode{1, &b}
	remove(&a, 2)
	fmt.Println(a)
	fmt.Println(a.Next)
	fmt.Println(a.Next.Next)
}

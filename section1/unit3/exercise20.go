// +build exercise20

package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func delete(h *listNode, k int) {
	count := 1
	tmpHead := &listNode{0, h}
	for ; tmpHead.Next != nil; tmpHead = tmpHead.Next {
		if count == k {
			tmpHead.Next = tmpHead.Next.Next
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
	delete(&a, 2)
	fmt.Println(a.Next)
}

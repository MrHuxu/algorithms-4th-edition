// +build exercise27

package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func max(h *listNode) int {
	result := h.Val
	for ; h != nil; h = h.Next {
		if h.Val > result {
			result = h.Val
		}
	}
	return result
}

func maxRec(n *listNode, v int) int {
	if n == nil {
		return v
	}
	if n.Val > v {
		v = n.Val
	}
	return maxRec(n.Next, v)
}

func main() {
	f := listNode{1, nil}
	e := listNode{2, &f}
	d := listNode{6, &e}
	c := listNode{5, &d}
	b := listNode{3, &c}
	a := listNode{4, &b}
	fmt.Println(max(&a))
	fmt.Println(maxRec(&a, a.Val))
}

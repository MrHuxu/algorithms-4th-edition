// +build weightedQuickUnionUf

package main

import "fmt"

type weightedQuickUnionUf struct {
	count int
	id    []int
	sz    []int
}

func newUnion(N int) *weightedQuickUnionUf {
	tmp := weightedQuickUnionUf{}
	tmp.count = N
	tmp.id = make([]int, N)
	tmp.sz = make([]int, N)
	for i := 0; i < N; i++ {
		tmp.id[i] = i
		tmp.sz[i] = 1
	}
	return &tmp
}

func (w *weightedQuickUnionUf) connected(p int, q int) bool {
	return w.find(p) == w.find(q)
}

func (w *weightedQuickUnionUf) find(p int) int {
	for p != w.id[p] {
		p = w.id[p]
	}
	return p
}

func (w *weightedQuickUnionUf) union(p int, q int) {
	i := w.find(p)
	j := w.find(q)
	if i == j {
		return
	}
	if w.sz[i] < w.sz[j] {
		w.id[i] = j
		w.sz[j] += w.sz[i]
	} else {
		w.id[j] = i
		w.sz[i] += w.sz[j]
	}
	w.count++
}

func main() {
	w := newUnion(10)
	w.union(4, 5)
	fmt.Println(w)
}

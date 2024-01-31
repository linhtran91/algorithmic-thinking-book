package main

import "fmt"

func main() {
	uf := NewQuickUnion(10)
	// 1-2-5-6-7 3-8-9 4
	uf.union(1, 2)
	uf.union(2, 5)
	uf.union(5, 6)
	uf.union(6, 7)
	uf.union(3, 8)
	uf.union(8, 9)

	fmt.Println(uf.connected(1, 5)) // true
	fmt.Println(uf.connected(5, 7)) // true
	fmt.Println(uf.connected(4, 9)) // false
	// 1-2-5-6-7 3-8-9-4
	uf.union(9, 4)
	fmt.Println(uf.connected(4, 9)) // true
}

type QuickUnion struct {
	root []int
}

func NewQuickUnion(size int) QuickUnion {
	root := make([]int, size)
	for i := 0; i < size; i++ {
		root[i] = i
	}
	return QuickUnion{root: root}
}

func (u *QuickUnion) find(x int) int {
	for x != u.root[x] {
		x = u.root[x]
	}
	return u.root[x]
}

func (u *QuickUnion) union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX != rootY {
		u.root[rootY] = rootX
	}
}

func (u *QuickUnion) connected(x, y int) bool {
	return u.find(x) == u.find(y)
}

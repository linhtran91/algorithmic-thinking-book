package main

import (
	"fmt"
)

func main() {
	uf := NewPathCompression(10)
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

type PathCompression struct {
	root []int
}

func NewPathCompression(size int) PathCompression {
	root := make([]int, size)
	for i := 0; i < size; i++ {
		root[i] = i
	}
	return PathCompression{root: root}
}

func (u *PathCompression) find(x int) int {
	if x == u.root[x] {
		return x
	}
	u.root[x] = u.find(u.root[x])
	return u.root[x]
}

func (u *PathCompression) union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}
	for i := 0; i < len(u.root); i++ {
		if u.root[i] == rootY {
			u.root[i] = rootX
		}
	}
}

func (u *PathCompression) connected(x, y int) bool {
	return u.find(x) == u.find(y)
}

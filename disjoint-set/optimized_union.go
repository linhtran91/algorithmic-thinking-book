package main

import "fmt"

func main() {
	uf := NewOptimizedUnion(10)
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

type OptimizedUnion struct {
	root []int
	rank []int
}

func NewOptimizedUnion(size int) OptimizedUnion {
	root := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		rank[i] = 1
		root[i] = i
	}
	return OptimizedUnion{rank: rank, root: root}
}

func (u *OptimizedUnion) find(x int) int {
	if x == u.root[x] {
		return x
	}
	u.root[x] = u.find(u.root[x])
	return u.root[x]
}

func (u *OptimizedUnion) union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}
	if u.rank[rootX] > u.rank[rootY] {
		u.root[rootY] = rootX
	} else if u.rank[rootX] < u.rank[rootY] {
		u.root[rootX] = rootY
	} else {
		u.root[rootY] = rootX
		u.rank[rootX] += 1
	}
}

func (u *OptimizedUnion) connected(x, y int) bool {
	return u.find(x) == u.find(y)
}

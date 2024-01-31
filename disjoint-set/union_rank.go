package main

import "fmt"

func main() {
	uf := NewRankUnion(10)
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

	// uf.union(0, 2)
	// uf.union(0, 1)
	// uf.union(0, 3)
	// uf.union(0, 4)
	// uf.union(5, 6)
	// uf.union(0, 5)
	// uf.union(6, 7)
}

type RankUnion struct {
	root []int
	rank []int
}

func NewRankUnion(size int) RankUnion {
	root := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		root[i] = i
		rank[i] = 1
	}
	return RankUnion{root: root, rank: rank}
}

func (u *RankUnion) find(x int) int {
	for x != u.root[x] {
		x = u.root[x]
	}
	return u.root[x]
}

func (u *RankUnion) union(x, y int) {
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
	fmt.Println(x, y)
	fmt.Println(u.root, u.rank)
}

func (u *RankUnion) connected(x, y int) bool {
	return u.find(x) == u.find(y)
}

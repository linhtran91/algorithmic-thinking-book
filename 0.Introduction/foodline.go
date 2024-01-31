package main

import "fmt"

// problem: https://dmoj.ca/problem/lkp18c2p1

func main() {
	nums := []int{3, 2, 5}
	showFoodLine(nums, 4)
}

func showFoodLine(nums []int, k int) {
	for i := 0; i < k; i++ {
		p := getMin(nums)
		fmt.Println(nums[p])
		nums[p] += 1
	}
}

func getMin(nums []int) int {
	min, p := nums[0], 0
	for i, v := range nums {
		if v < min {
			v = min
			p = i
		}
	}
	return p
}

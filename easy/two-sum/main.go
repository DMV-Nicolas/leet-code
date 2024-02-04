package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		aim := target - v

		if aimIndex, ok := m[aim]; ok {
			return []int{aimIndex, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := twoSum(nums, target)
	fmt.Println(sum)
}

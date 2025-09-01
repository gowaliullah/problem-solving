package main

import "fmt"

func twoSum(nums []int, target int) []int { // O(n^2)
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int { // O(n)
	m := make(map[int]int)

	for i, num := range nums {
		if val, ok := m[target-num]; ok {
			fmt.Println([]int{val, i})
			return []int{val, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
	twoSum2(nums, target)
}

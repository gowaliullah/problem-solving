package main

// 1464. Maximum Product of Two Elements in an Array

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	sort.Ints(nums[:])

	max1 := nums[len(nums)-2]
	max2 := nums[len(nums)-1]

	return (max1 - 1) * (max2 - 1)

}

func main() {
	arr := [4]int{3, 4, 5, 2}

	res := maxProduct(arr[:])

	fmt.Println(res)

}

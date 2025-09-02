package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	newLength := removeElement(nums, val)
	fmt.Println("New length:", newLength)
	fmt.Println("Modified array:", nums[:newLength])
}

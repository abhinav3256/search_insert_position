package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	result := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == target || nums[i] > target {
			result = i
			break
		}
	}
	len := len(nums)
	last_index := len - 1

	if nums[last_index] < target {
		result = len
	}

	// fmt.Println(len)
	// fmt.Println(last_index)

	fmt.Println(result)

}

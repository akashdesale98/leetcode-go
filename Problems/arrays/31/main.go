package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println("nums", nums)
}

func nextPermutation(nums []int) {
	index := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}

	if index != -1 {
		fmt.Println("if")
		for i := len(nums) - 1; i >= index+1; i-- {
			if nums[index] < nums[i] {
				nums[index], nums[i] = nums[i], nums[index]
				break
			}
		}

		reverse(nums[index+1:])
	} else {
		fmt.Println("else")
		reverse(nums)
	}

}

func reverse(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return nums
}

// https://leetcode.com/problems/move-zeroes/
package main

import "fmt"

func main() {
	nums := []int{1}
	moveZeroesOpt(nums)
	fmt.Println("nums", nums)
}

func moveZeroes(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	temp := []int{}
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			temp = append(temp, nums[i])
		}
	}
	copy(nums, temp)

	for i := len(temp); i < n; i++ {
		nums[i] = 0
	}
}

func moveZeroesOpt(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	j := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j = i
			break
		}
	}

	if j == -1 {
		return
	}

	for i := j + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

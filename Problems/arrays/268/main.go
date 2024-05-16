// https://leetcode.com/problems/missing-number
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	if len(nums) == 1 && nums[0] != 0 {
		return 0
	}

	sort.Ints(nums)

	if len(nums) == 1 && nums[0] != 0 {
		return 0
	}

	j := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != j {
			break
		}
		j++
	}

	return j
}

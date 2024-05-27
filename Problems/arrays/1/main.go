// https://leetcode.com/problems/two-sum/
package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	pair, result := map[int]int{}, []int{}

	for i := 0; i < len(nums); i++ {
		if index, ok := pair[target-nums[i]]; ok {
			return []int{index, i}
		}
		pair[nums[i]] = i
	}

	return result
}

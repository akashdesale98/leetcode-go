// https://leetcode.com/problems/rearrange-array-elements-by-sign/
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, -1}
	fmt.Println(rearrangeArray(nums))
}

func rearrangeArray(nums []int) []int {
	positive, negative := []int{}, []int{}

	for i := 0; i < len(nums)/2; i++ {
		if nums[i] < 0 {
			negative = append(negative, nums[i])
		} else {
			positive = append(positive, nums[i])
		}
	}

	for i := 0; i < len(negative)/2; i++ {
		nums[2*i] = positive[i]
		nums[2*i+1] = negative[i]
	}

	return nums
}

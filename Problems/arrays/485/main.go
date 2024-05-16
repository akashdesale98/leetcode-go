// https://leetcode.com/problems/max-consecutive-ones
package main

import "fmt"

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))

}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	intersection := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			intersection = append(intersection, count)
			count = 0
		}
	}

	intersection = append(intersection, count)

	count = 0
	for i := 0; i < len(intersection); i++ {
		if count < intersection[i] {
			count = intersection[i]
		}
	}

	return count
}

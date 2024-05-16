// https://leetcode.com/problems/single-number/
package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {

	appearances := map[int]int{}
	for i := 0; i < len(nums); i++ {
		appearances[nums[i]]++
	}

	for k, v := range appearances {
		if v == 1 {
			return k
		}
	}

	return 0
}

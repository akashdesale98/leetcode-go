// https://leetcode.com/problems/rotate-array
package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotateRight(nums, 3)
}

func rotateRight(nums []int, k int) {

	rotateLeft(nums, len(nums)-k)
}

func rotateLeft(nums []int, k int) {

	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}

	temp := make([]int, k, k)
	copy(temp, nums[:k])

	for i := k; i < n; i++ {
		nums[i-k] = nums[i]
	}

	for i := n - k; i < n; i++ {
		nums[i] = temp[i-(n-k)]
	}

	fmt.Println("nums", nums)
}

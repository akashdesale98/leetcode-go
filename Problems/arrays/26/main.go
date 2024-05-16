package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
import "fmt"

func main() {
	nums := []int{0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {
	j := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i+1]
			j++
		}
	}

	return j
}

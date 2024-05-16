package main

// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/submissions/1258630400/
import "fmt"

func main() {
	nums := []int{2, 1}

	fmt.Println(check(nums))

}

func check(nums []int) bool {

	counter := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			counter++
			if counter > 1 || nums[0] < nums[len(nums)-1] {
				return false
			}
		}
	}

	return true
}

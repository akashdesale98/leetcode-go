// https://leetcode.com/problems/intersection-of-two-arrays
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	intersection := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			if len(intersection) == 0 {
				intersection = append(intersection, nums1[i])
			}
			if len(intersection) > 0 && intersection[len(intersection)-1] != nums1[i] {
				intersection = append(intersection, nums1[i])
			}
			i++
			j++
		}
	}

	return intersection
}

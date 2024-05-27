// https://leetcode.com/problems/search-a-2d-matrix
package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	low, high := 0, n*m-1
	for low <= high {
		mid := (low + high) / 2
		row, col := mid/m, mid%m
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

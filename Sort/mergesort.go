package main

import (
	"fmt"
)

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	done := make(chan bool)
	mid := len(data) / 2
	var left []int

	go func() {
		left = mergeSort(data[:mid])
		done <- true
	}()
	right := mergeSort(data[mid:])
	<-done
	return merge(left, right)

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{7, 12, 11, 13, 13, 6, 7}
	fmt.Printf("%v\n%v\n", arr, mergeSort(arr))

}

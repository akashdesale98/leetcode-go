package main

func main() {
	nums := []int{5, 1, 2, 3, 4}

	// selectionSort(nums)
	// bubbleSort(nums)
	// insertionSort(nums)
	mergeSort(nums)
}

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i <= n-2; i++ {
		min := i
		for j := i; j <= n-1; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}

func bubbleSort(nums []int) {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func insertionSort(nums []int) {
	n := len(nums)

	for i := 0; i <= n-1; i++ {
		j := i
		for j > 0 {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			j--
		}
	}
}

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

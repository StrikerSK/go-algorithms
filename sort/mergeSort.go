package sort

import "fmt"

// MergeSort - recursively divide the array into two halves, sort each half and then merging them
func MergeSort() {
	input := []int{56, 123, 2, 78, 15, 79, 35, 89, 20, 54}
	fmt.Println("Input array: ", input)
	fmt.Println("Output array: ", mergeSort(input))
}

func mergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	} else {
		middle := len(input) / 2
		leftArray := mergeSort(input[0:middle])
		rightArray := mergeSort(input[middle:])

		return mergeSlices(leftArray, rightArray)
	}
}

func mergeSlices(slice1, slice2 []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			result = append(result, slice1[i])
			i++
		} else {
			result = append(result, slice2[j])
			j++
		}
	}

	for i < len(slice1) {
		result = append(result, slice1[i])
		i++
	}

	for j < len(slice2) {
		result = append(result, slice2[j])
		j++
	}

	return result
}

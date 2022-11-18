package sort

import "fmt"

// QuickSort -
func QuickSort() {
	input := []int{56, 123, 2, 78, 15, 79, 35, 89, 20, 54}
	fmt.Println("Input array: ", input)
	fmt.Println("Output array: ", quickSort(input, 0, len(input)-1))
}

func quickSort(slice []int, low, high int) []int {
	if low < high {
		p := pivotUtil(slice, low, high)
		quickSort(slice, low, p-1)
		quickSort(slice, p+1, high)
	}

	return slice
}

func pivotUtil(slice []int, low, high int) int {
	pivot := slice[high]
	i := low

	for j := low; j < high; j++ {
		if slice[j] < pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}
	slice[i], slice[high] = slice[high], slice[i]
	return i
}

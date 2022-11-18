package sort

import "fmt"

// QuickSort -
func QuickSort(input []int) {
	var inputCopy []int
	inputCopy = append(inputCopy, input...)
	fmt.Println("Quick sort - input array: ", inputCopy)
	fmt.Println("Quick sort - output array: ", quickSort(inputCopy, 0, len(inputCopy)-1))
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

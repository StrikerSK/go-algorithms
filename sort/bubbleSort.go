package sort

import "fmt"

// BubbleSort - compares value with rest of the array's values and swaps the value with the lowest value
func BubbleSort(input []int) {
	var inputCopy []int
	inputCopy = append(inputCopy, input...)
	fmt.Println("Bubble sort - input array: ", inputCopy)
	fmt.Println("Bubble sort - output array: ", bubbleSorter(inputCopy))
}

func bubbleSorter(input []int) []int {
	var isSwapped bool

	for i := 0; i < len(input)-1; i++ {
		isSwapped = false

		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}

	return input
}

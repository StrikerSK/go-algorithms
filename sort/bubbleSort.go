package sort

import "fmt"

// BubbleSort - compares value with rest of the array's values and swaps the value with the lowest value
func BubbleSort() {
	input := []int{56, 123, 2, 78, 15, 79, 35, 89, 20, 54}
	fmt.Println("Input array: ", input)
	fmt.Println("Output array: ", bubbleSorter(input))
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

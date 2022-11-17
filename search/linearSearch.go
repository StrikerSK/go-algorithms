package search

import "fmt"

// LinearSearch - loops through the array and compares each element with the search value
func LinearSearch() {
	searchedValue := 15
	input := []int{56, 123, 2, 78, 15, 79, 35, 89, 20, 54}

	for i := 0; i < len(input); i++ {
		if input[i] == searchedValue {
			fmt.Printf("Value %d present at index %d\n", searchedValue, i)
		}
	}
}

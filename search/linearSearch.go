package search

import "fmt"

// LinearSearch - loops through the array and compares each element with the search value
func LinearSearch(searchValue int, input []int) {
	for i := 0; i < len(input); i++ {
		if input[i] == searchValue {
			fmt.Printf("Value %d present at index %d\n", searchValue, i)
		}
	}
}

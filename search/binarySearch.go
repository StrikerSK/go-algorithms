package search

import "fmt"

// BinarySearch - should be done on sorted array
func BinarySearch(searchValue int, input []int) {
	firstIndex := 0
	lastIndex := len(input) - 1
	middleIndex := (firstIndex + lastIndex) / 2

	for input[middleIndex] != searchValue && firstIndex < lastIndex {
		if input[middleIndex] > searchValue {
			lastIndex = middleIndex - 1
		} else {
			firstIndex = middleIndex + 1
		}
		middleIndex = (firstIndex + lastIndex) / 2
	}

	fmt.Printf("Value %d present at index %d\n", searchValue, middleIndex)
}

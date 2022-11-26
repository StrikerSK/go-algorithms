package main

import (
	"fmt"
	"go-sorting/search"
	"go-sorting/sort"
)

func main() {
	sectionDivider("Sorting")
	input := []int{56, 123, 2, 78, 15, 79, 35, 89, 20, 54}
	sort.MergeSort(input)
	sort.QuickSort(input)
	sort.BubbleSort(input)
	sort.RadixSort(input)

	sectionDivider("Searching")
	sortedInput := []int{2, 15, 20, 35, 54, 56, 78, 79, 89, 123}
	searchValue := 15
	search.LinearSearch(searchValue, sortedInput)
	search.BinarySearch(searchValue, sortedInput)
}

func sectionDivider(value string) {
	fmt.Printf("----------------------------- %s -----------------------------\n", value)
}

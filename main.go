package main

import (
	"go-sorting/sort"
)

func main() {
	input := []int{56, 123, 2, 78, 15, 79, 35, 89, 20, 54}
	sort.MergeSort(input)
	sort.QuickSort(input)
	sort.BubbleSort(input)
}

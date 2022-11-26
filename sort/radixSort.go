package sort

import (
	"fmt"
	"math"
	"strconv"
)

func RadixSort(input []int) {
	var inputCopy []int
	inputCopy = append(inputCopy, input...)
	fmt.Println("Radix sort - input array: ", inputCopy)
	fmt.Println("Radix sort - output array: ", sort(inputCopy))
}

func sort(input []int) []int {
	maxDigitCount := maxDigit(input)
	var output []int

	for i := 0; i < maxDigitCount; i++ {
		var digitBuckets [10][]int

		for j := 0; j < len(input); j++ {
			lastDigit := getDigit(input[j], i)
			digitBuckets[lastDigit] = append(digitBuckets[lastDigit], input[j])
		}
		fmt.Println("console")
		output = append(output, digitBuckets[0]...)
	}

	return input
}

func maxDigit(slice []int) int {
	maxNumber := 0
	for _, number := range slice {
		if number > maxNumber {
			maxNumber = int(math.Max(float64(maxNumber), float64(countDigits(number))))
		}
	}

	return maxNumber
}

func getDigit(number, index int) int {
	stringNumber := strconv.Itoa(int(math.Abs(float64(number))))
	currentIndex := len(stringNumber) - index - 1

	stringToNumber, err := strconv.Atoi(string([]rune(stringNumber)[currentIndex]))
	if err != nil {
		return 0
	}

	return stringToNumber
}

func countDigits(number int) int {
	if number == 0 {
		return 1
	}

	return int(math.Floor(math.Log10(float64(len(strconv.Itoa(number))))) + 1)
}

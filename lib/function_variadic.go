package lib

import "fmt"

func FunctionVariadic() {
	var sum int = sum(1, 2, 3, 4, 5)
	fmt.Println("Hasil Penjumlahan dari function sum: ", sum)

	// pengisian variadic dengan data slice
	var unsortedArray = []int{5, 2, 7, 3, 8, 9, 1}
	var sortedArray = sort("desc", unsortedArray...)
	fmt.Println("Sorted Array:", sortedArray)
}

func sum(numbersArray ...int) int {
	var output int = 0
	for _, number := range numbersArray {
		output += number
	}

	return output
}

// parameter variadic harus selalu di bagian belakang
func sort(mode string, numbers ...int) []int {
	switch mode {
	case "asc": {
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				if numbers[j] < numbers[i] {
					var temp int = numbers[j]
					numbers[j] = numbers[i]
					numbers[i] = temp
				}
			}
		}
	}
	case "desc": {
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				if numbers[j] > numbers[i] {
					var temp int = numbers[j]
					numbers[j] = numbers[i]
					numbers[i] = temp
				}
			}
		}
	}
	default: {
		fmt.Println("Invalid sorting mode")
		return numbers
	}
	}

	return numbers
}

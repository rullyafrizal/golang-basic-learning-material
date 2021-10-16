package lib

import (
	"fmt"
	"strings"
)

func FunctionClosure() {
	// Closure == Anonymous Function == Function inside variable

	// Example 1
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %d\nmin: %d\nmax: %d\n", numbers, min, max)


	// Example 2
	var palindrome = func (word string) (bool) {
		var arrWord = strings.Split(word, "")
		var reversed string

		for i := len(arrWord) - 1; i >= 0; i-- {
			reversed += arrWord[i]
		}

		if reversed == word {
			return true
		} else {
			return false
		}
	}

	word := "rully"
	// %v digunakan untuk segala jenis tipe data
	fmt.Printf("Word %s is a palindrome? %v\n", word, palindrome(word))


	// IIFE = Immediately Invoked Function Expression
	// Function closure yang langsung di invoke
	numbers = []int{2, 3, 0, 4, 3, 2, 1, 4, 2, 0, 3}
	var newNumbers = func(min int) []int {
		var r[] int

		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}

		return r
	}(2)
	fmt.Println("Original Number : ", numbers)
	fmt.Println("Filtered Numbers : ", newNumbers)


	max = 3
	numbers = []int{2, 3, 0, 4, 3, 2, 1, 4, 2, 0, 3}

	var count, getNumbers = findMax(max, numbers...)
	var theNumbers = getNumbers()

	fmt.Println("number : ", numbers)
	fmt.Printf("find : %d\n\n", max)

	fmt.Println("found : ", count)
	fmt.Println("value : ", theNumbers)
}

// Closure sebagai return value
func findMax(max int, numbers ...int) (int, func() []int) {
	var res[] int

	for _, number := range numbers {
		if number <= max {
			res = append(res, number)
		}
	}

	//return len(res), func() []int {
	//	return res
	//}

	// atau boleh di simpan di variabel dulu
	var getNumbers = func() []int {
		return res
	}

	return len(res), getNumbers
}

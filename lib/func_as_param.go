package lib

import (
	"fmt"
	"strings"
)

func FunctionAsParameter() {
	// Function as a Parameter = Callback Function
	var data = []string{"Rully", "Afrizal", "Alwin"}
	var dataContainsA = filter(data, func(item string) bool {
		return strings.Contains(item, "a") || strings.Contains(item, "A")
	})

	fmt.Println("data asli : ", data)
	fmt.Println("filter harus ada huruf a :", dataContainsA)


	// Example 2
	var dataLength5 = filterAlias(data, func(item string) bool {
		return len(item) == 5
	})
	fmt.Println("\ndata asli : ", data)
	fmt.Println("filter harus ada huruf a :", dataLength5)
}

func filter(data[] string, callback func(string) bool) []string {
	var result []string

	for _, item := range data {
		if filtered := callback(item); filtered {
			result = append(result, item)
		}
	}

	return result
}


// type alias sebagai skema closure
func filterAlias(data[] string, callback FilterCallback) []string {
	var result []string

	for _, item := range data {
		if filtered := callback(item); filtered {
			result = append(result, item)
		}
	}

	return result
}


// alias
type FilterCallback func(string) bool

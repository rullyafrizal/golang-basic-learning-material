package concurrent

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var numbers = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	fmt.Println("numbers", numbers)

	var ch1 = make(chan float64)
	go getAvg(numbers, ch1)

	var ch2 = make(chan float64)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %.2f \n", max)
		}
	}
}

func getAvg(numbers []int, channel chan float64) {
	var sum = 0

	for _, number := range numbers {
		sum += number
	}

	channel <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, channel chan float64) {
	var max = 0

	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	channel <- float64(max)
}

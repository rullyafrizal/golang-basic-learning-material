package lib

import (
	"fmt"
	"math"
)

func FunctionMultipleReturn() {
	var area, circumference = circleCalculation(10)
	fmt.Printf("Luas Lingkaran: %.3f, Keliling: %.3f\n", area, circumference)

	area, circumference = squareCalculation(5)
	fmt.Printf("Luas Persegi: %.1f, keliling: %.1f\n", area, circumference)
}

func circleCalculation (d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d / 2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

// pre-defined return value
func squareCalculation (s float64) (area float64, circumference float64) {
	area = math.Pow(s, 2)
	circumference = s * 4

	return
}

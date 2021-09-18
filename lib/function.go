package lib

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Function() {
	var names = []string{"Rully", "Afrizal", "Alwin"}
	sayHello("Selamat Pagi", names)
	fmt.Printf("Hasil Penjumlahan : %d\n", simpleSumFunction(5, 10))
	rand.Seed(time.Now().Unix())
	fmt.Printf("Angka Random: %d\n", randomWithRange(2, 100))
	divideNumber(5, 2)
}

func sayHello (message string, arr []string) {
	var name = strings.Join(arr, " ")
	fmt.Printf("%s %s\n", message, name)
}

func simpleSumFunction(a, b int) int {
	return a + b
}

func randomWithRange(min, max int) int {
	return rand.Int() % (max - min + 1) + min
}

func divideNumber(m, n float32) {
	if n == 0 {
		fmt.Printf("Invalid number, %.2f cannot be divided by %.2f\n", m, n)
		return
	}

	var output = m / n
	fmt.Printf("%.2f / %.2f = %.2f\n", m, n, output)
}

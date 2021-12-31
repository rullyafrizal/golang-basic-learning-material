package lib

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	// error merupakan sebuah tipe, error memiliki 1 buah property berupa method Error()
	// method Error() return value-nya adalah pesan error (string), bisa nil

	var input string = "rok" // error karena bukan number

	var num int
	var err error

	// strconv berguna untuk konversi string ke numerik
	// return value ada 2, yaitu number dan error
	// apabila data bisa di-konversi, maka return number, jika tidak maka error
	num, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(num, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}

func TestCustomError(t *testing.T) {
	name := "   "
	name2 := "Rully  "

	// error
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
	}

	// success
	if valid, err := validate(name2); valid {
		fmt.Println("halo", name2)
	} else {
		fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Input cannot be empty")
	}

	return true, nil
}

func TestPanic(t *testing.T) {
	// panic digunakan untuk menampilkan stack trace error
	// sekaligus menghentikan flow Goroutine

	var name string = "   "

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println("Error bro")
		panic(err.Error())
		fmt.Println("end")
	}
}

func TestRecover(t *testing.T) {
	// recover digunakan untuk handle panic error.
	// recover take-over Goroutine ketika panic sedang terjadi, sehingga pesan panic tidak muncul
	defer catch()

	var name string = "   "

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println("Error bro")
		panic(err.Error())
		fmt.Println("end")
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func TestRecoverIife(t *testing.T) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Panic occured", r)
	//	} else {
	//		fmt.Println("Application running perfectly")
	//	}
	//}()
	//
	//panic("Some error happen")

	// di atas adalah contoh penerapan recover pada IIFE
	// tetapi dalam real-world development,
	// ada kalanya recover tidak dibutuhkan dalam blok fungsi terluar, tetapi di dalam blok fungsi yang lebih spesifik

	data := []string{"Superman", "Batman", "Thanos", "Ironman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message: ", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()
		}()
	}

	panic("Some error happen")
}

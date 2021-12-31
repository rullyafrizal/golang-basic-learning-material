package lib

import (
	"fmt"
	"testing"
)

func TestFormatString(t *testing.T) {
	// %b -> digunakan untuk memformat data numerik menjadi bentuk string numerik berbasis 2 (binary)
	fmt.Printf("age : %b\n", data.age)

	// %c -> digunakan untuk format data numerik yang merupaka unicode
	// menjadi bentuk string karakter unicodenya
	fmt.Printf("%c \n", 1400) // n

	// %d -> digunakan untuk format data numerik ke string numerik basis 10
	fmt.Printf("age : %d \n", data.age)

	// %e atau %E -> digunakan untuk format data numerik desimal ke dalam bentuk notasi numerik standar
	// contoh : 1.825000e+02 ->yang artinya 1.825 * 10^2
	fmt.Printf("height : %e \n", data.height)

	// %f atau %F
	// berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan
	fmt.Printf("height: %.2f \n", data.height)

	// %g atau %G
	// berfungsi untuk memformat data numerik desimal besar
	fmt.Printf("%g \n", 0.2414235716561561)

	// %o
	// digunakan untuk memforat data numerik, menjadi bentuk string basis 8 (oktal)
	fmt.Printf("octal age : %o \n", data.age)

	// %p
	// digunakan untuk memformat data pointer, return alamat pointer referensi variabel-nya
	fmt.Printf("name pointer address : %p \n", &data.name)

	// lebih lengkap
	// kunjungi https://dasarpemrogramangolang.novalagung.com/A-string-format.html
}

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "Rully",
	height:      172.4,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"Swimming", "Coding"},
}

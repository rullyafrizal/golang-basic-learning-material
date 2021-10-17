package concurrent

import (
	f "fmt"
	"runtime"
)

func Goroutine() {
	// Goroutine mirip dengan thread thread, lebih pas disebut sebagai mini thread
	// Eksekusi Goroutine bersifat asynchronous
	// Goroutine merupakan salah satu bagian penting dalam concurrent programming di Go

	// proses yang akan dieksekusi goroutine harus dibungkus dengan function
	// tambahkan keyword go di depan proses yang akan dieksekusi

	// menentukan jumlah core processor yang akan digunakan untuk eksekusi
	runtime.GOMAXPROCS(4)

	// dengan goroutine
	go print(1000, "goroutine")

	// tanpa goroutine
	print(1000, "tanpa goroutine")

	var input string
	f.Scanln(&input)
}

func print(max int, message string) {
	for i := 1; i <= max; i++ {
		f.Println(message, "-->", i)
	}
}



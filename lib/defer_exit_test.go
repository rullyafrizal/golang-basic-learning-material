package lib

import (
	"fmt"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	// Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok function selesai
	defer fmt.Println("Ini defer, di-print terakhir")

	fmt.Println("bukan defer, di-print pertama")
}

func orderFood(menu string) {
	defer fmt.Println("Terimakasih, silahkan tunggu :)")

	switch menu {
	case "pizza":
		fmt.Println("Pilihan tepat!")
		fmt.Println("Pizza di tempat kami paling enak!")
	case "burger":
		fmt.Println("Yahh, stok burger kosong :(")
	}

	fmt.Println("Pesanan anda: ", menu)
}

func TestOrderFood(t *testing.T) {
	orderFood("pizza")
}

// Kombinasi defer dan IIFE
func TestDeferIIFE(t *testing.T) {
	num := 3

	if num == 3 {
		fmt.Println("Halo 1")
		//defer fmt.Println("halo 3") // defer akan tetap dieksekusi di-akhir, meskipun di dalam blok if

		// apabila ingin langsung dieksekusi di akhir blok if, harus dibungkus IIFE
		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")
}

func TestExit(t *testing.T) {
	// Exit digunakan untuk menghentikan program secara paksa pada saaat itu juga
	// semua statement setelah exit tidak akan dieksekusi, termasuk juga defer

	defer fmt.Println("Halo defer")
	os.Exit(1) // terdapat exit code status yang harus di-passing
	fmt.Println("Selamat datang!")
}

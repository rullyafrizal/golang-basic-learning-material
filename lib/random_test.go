package lib

// docs : https://dasarpemrogramangolang.novalagung.com/A-random.html

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomSeed(t *testing.T) {
	// Fungsi rand.Seed()digunakan untuk penentuan nilai seed.
	// Fungsi rand.Int() digunakan untuk generate angka random dalam bentuk numerik bertipe int.
	// Fungsi rand.Int() ini setiap kali dipanggil akan menghasilkan angka berbeda, tapi pasti hasilnya akan selalu tetap jika mengacu ke deret.
	// Angka random ke-1 akan selalu 5221277731205826435
	// Angka random ke-2 akan selalu 3852159813000522384
	// Angka random ke-3 akan selalu 8532807521486154107
	// Dan seterusnya ...
	rand.Seed(10)

	fmt.Printf("random ke-1 : %d\n", rand.Int())
	fmt.Printf("random ke-2 : %d\n", rand.Int())
	fmt.Printf("random ke-3 : %d\n", rand.Int())
}

func TestRandomUniqueSeed(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano()) // seed uniquer menggunakan unix nano dari waktu sekarang
	fmt.Printf("random ke-1 : %d\n", rand.Int())
	fmt.Printf("random ke-2 : %d\n", rand.Int())
	fmt.Printf("random ke-3 : %d\n", rand.Int())
}

func TestRandomString(t *testing.T) {
	fmt.Printf("random string : %s\n", randomString(5))
}

func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

package lib

import "fmt"

func Pointer() {
	// Pointer adalah reference atau alamat memory
	// Sebagai contoh {var x int = 4}
	// Maka Pointer adalah alamat memory di mana nilai 4 dari variabel x disimpan

	// Variabel pointer ditandai dengan tanda "*" (arterisk)
	//var number *int
	//var name *string

	// Poin Penting:
	// 1. Variabel biasa bisa diambil nilai pointernya, cukup tambahkan "&" tepat sebelum nama variabel, metode ini dinamakan referencing
	// 2. Nilai asli variabel pointer bisa diambil, cukup tambahkan "*" tepat sebelum nama variabel, metode ini dinamakan dereferencing

	// Example
	var numberA int = 4
	var numberB *int = &numberA
	// variabel number b adalah variabel pointer,
	// yang isinya adalah variabel biasa yang diambil nilai pointernya menggunakan &

	fmt.Println("number A : ", numberA)
	fmt.Println("Address number A : ", numberB)

	fmt.Println("number B : ", *numberB)
	fmt.Println("Address number B : ", numberB)


	// Efek perubahan nilai Pointer
	// ketika salah satu variabel pointer di ubah nilainya, sedangkan ada
	// variabel lain yang memiliki referensi memori yang sama, maka
	// nilai variabel lain tersebut juga akan berubah selama memory reference yang dimiliki sama
	var numberC int = 5
	var numberD *int = &numberC
	fmt.Println("number C : ", numberC)
	fmt.Println("Address number C : ", numberD)
	fmt.Println("number D : ", *numberD)
	fmt.Println("Address number D : ", numberD)

	numberC = 6

	// dalam hal ini nilai variabel D akan berubah ke 6 juga
	fmt.Println("number C : ", numberC)
	fmt.Println("Address number C : ", numberD)
	fmt.Println("number D : ", *numberD)
	fmt.Println("Address number D : ", numberD)


	// Parameter Pointer
	var number = 4
	fmt.Println("before : ", number) // 4

	change(&number, 10)
	fmt.Println("after : ", number) // 10
}

// Parameter Pointer
func change(original *int, value int) {
	*original = value
}

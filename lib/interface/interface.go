package _interface

import "fmt"

// Interface adalah kumpulan definisi method yang tidak memiliki isi, yang dibungkus
// dengan nama tertentu

// Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah nil.

// variabel interface hanya mampu menampung objek yang memiliki semua method yang terdefinisi di interface
// konsepnya mirip seperti interface di bahasa OOP lain
// class yang implement suatu interface maka harus implement semua method yang ada di interface tersebut

type Hitung interface {
	luas() float64
	keliling() float64
}

func InterfaceHitung() {
	var bangunDatar Hitung

	bangunDatar = persegi{10.0}
	fmt.Println("====Persegi====")
	fmt.Println("Luas :", bangunDatar.luas())
	fmt.Println("Keliling :", bangunDatar.keliling())

	bangunDatar = lingkaran{10.0}
	fmt.Println("\n====Lingkaran====")
	fmt.Println("Luas :", bangunDatar.luas())
	fmt.Println("Keliling :", bangunDatar.keliling())
	fmt.Println("Jari-jari :", bangunDatar.(lingkaran).jariJari()) // casting dulu karena interface Hitung tidak memiliki function jariJari()
	// atau menggunakan cara satu ini juga bisa
	var bangunLingkaran lingkaran = bangunDatar.(lingkaran)
	fmt.Println("Jari-jari :", bangunLingkaran.jariJari())

}

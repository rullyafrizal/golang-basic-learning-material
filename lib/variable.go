package lib

import "fmt"

func Variable() {
	/**
	* skema penggunaan keyword var
	* var <nama-variabel> <tipe-data>
	* var <nama-variabel> <tipe-data> = nilai
	 */
	var firstName string = "rully"

	var lastName string
	lastName = "afrizal"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	// sama dengan; jika menggunakan (,) maka tak perlu spasi
	fmt.Println("halo", firstName, lastName + "!")
	fmt.Println("halo " + firstName + " " + lastName + "!")

	/**
	* Deklarasi variabel tanpa tipe data (type inference)
	* <nama-variabel> := value
	 */

	var merekMobil string = "mercedes"
	kecepatan := 40

	fmt.Printf("mobil %s melaju dengan kecepatan %d km/h \n", merekMobil, kecepatan)

	/**
	* Deklarasi multi variabel
	 */
	//var first, second, third string
	//first, second, third = "satu", "dua", "tiga"

	//var fourth, fifth, sixth string = "empat", "lima", "enam"

	//multi variabel declaration with type inference
	//seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// type inference bisa digunakan untuk tipe berbeda
	//one, isTrue, twoPointTwo, say := 1, true, 2.2, "hello"

	/**
	* Variabel Underscore (_)
	* variabel ini tidak untuk digunakan, isi nya tidak akan ditampilkan
	* data yang sudah masuk ke variabel tersebut akan hilang
	 */
	//satu, _ := 1, "variabel underscore"

	/**
	* Deklarasi variabel menggunakan keyword new
	* - berfungsi untuk membuat variabel pointer
	* - jika ditampilkan akan muncul memory address dari variabel tersebut
	* - jika ingin menapilkan nilainya menggunakan (*) yang nanti akan dibahas
	 */
	//name := new(string)

	/**
	* Deklarasi variabel menggunakan keyword make
	 */
}

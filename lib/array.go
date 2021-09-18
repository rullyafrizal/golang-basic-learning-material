package lib

import "fmt"

func Array() {
	/**
	* Deklarasi array di GO
	* 1. var <nama-variabel> [jumlah-elemen]<tipe-data>
	* 2. var <nama-variabel> = [jumlah-elemen]<tipe-data>{value}
	 */
	var names [4]string
	names[0] = "Rully"
	names[1] = "Afrizal"
	names[2] = "Alwin"
	names[3] = "Ardelia"
	//names[4] = "Pramesti" //error karena melebihi alokasi awal

	fmt.Println(names[0], names[1], names[2], names[3])

	var cars = [4]string{"mercedes", "bmw", "chevrolet", "honda"} //horizontal
	fmt.Println(cars[0], cars[1], cars[2], cars[3])

	var primes = [4]uint8{
		2,
		3,
		5,
		7, //wajib trailing-comma
	}
	fmt.Println(primes[0], primes[1], primes[2], primes[3])


	/**
	* Array multidimensi
	 */
	var randomNumbers = [2][3]int{[3]int{5, 4, 6}, [3]int{10, 4, 3}}
	var randomNumbers2 = [2][3]int{{5, 4, 6}, {10, 4, 3}}
	fmt.Println(randomNumbers)
	fmt.Println(randomNumbers2)

	// Looping Array
	var fruits = [4]string{"banana", "apple", "pineapple", "orange"}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Fruits index %d: %s\n", i, fruits[i])
	} //menggunakan for biasa

	for i, fruit := range fruits {
		fmt.Printf("Fruits index %d: %s\n", i, fruit)
	} //menggunakan for-range (mirip for-each)

	// penggunana variabel underscore (_) dalam for-range
	//for i, fruit := range fruits {
	//	fmt.Printf("nama buah: %s\n", fruit)
	//} //error karena variabel i tidak dipakai

	for _, fruit := range fruits {
		fmt.Printf("nama buah: %s\n", fruit)
	} //tidak error karena memakai variabel underscore

	/**
	* Deklarasi array dengan keyword "make"
	* var <nama-variabel> = make([]<tipe-data>, <jumlah-elemen>)
	 */

	var animals = make([]string, 2)
	animals[0] = "Elephant"
	animals[1] = "Lion"
	for _, animal := range animals {
		fmt.Printf("nama hewan : %s\n", animal)
	}
}

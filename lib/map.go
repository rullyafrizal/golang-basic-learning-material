package lib

import "fmt"

func Map() {
	/**
	* map mirip object di js, atau array associative di php
	* Deklarasi map:
	* var <nama-variabel> = map[<tipe-data-key>]<tipe-data-value>{elemen}
	 */

	var jumlahChicken =  map[string]int{"kfc": 44, "mcd": 50, "aw": 45}
	var cars map[string]int
	//cars["matic"] = 1 //error karena belum diinisialisasi pakai curly bracket
	cars = map[string]int{
		"matic": 5,
		"manual": 10,
	}

	//accessing map, sama seperti di PL lain
	fmt.Println(jumlahChicken["kfc"])
	fmt.Println(cars["manual"])

	// for-range dalam mengakses map
	for key, val := range cars {
		fmt.Println(key + ":", val)
	}

	//menghapus item map
	var chickens = map[string]int{
		"januari": 50,
		"februari": 53,
		"maret": 51,
	}

	fmt.Println(len(chickens))
	fmt.Println(chickens)

	//hapus menggunakan delete
	delete(chickens, "maret")

	fmt.Println(len(chickens))
	fmt.Println(chickens)


	// mendeteksi key existing
	var value, isExist = chickens["januari"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("doesn't exist")
	}

	// kombinasi slice dan map
	var movies = []map[string]string{
		{"title": "Avengers", "release_date": "2 January 2020"},
		{"title": "Iron Man", "release_date": "2 January 2020"},
		{"title": "Black Panther", "release_date": "2 January 2020"},
		{"title": "Fantastic Four", "release_date": "2 January 2020"},
	}

	for _, movie := range movies {
		fmt.Printf("%s released at %s\n", movie["title"], movie["release_date"])
	}
}

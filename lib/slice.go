package lib

import "fmt"

func Slice() {
	/**
	* Slice adalah reference elemen array
	* karena merupaka data reference, perubahan pada slice akan berdampak pada slice lain yang
	* memory addressnya sama
	 */

	var fruits = []string{"apple", "grape", "banana"} //tidak perlu menulis ukuran elemen
	fmt.Println(fruits)

	var cars = []string{"volvo", "bmw"} //slice
	var animals = [2]string{"monkey", "ant"} //array
	var humans = [...]string{"rully", "afrizal", "alwin"} //array
	fmt.Println(cars, animals)
	fmt.Scanln(humans)

	/**
	* Hubungan slice dan array
	* slice bisa dibentuk dari array yang sudah didefinisikan
	* karena slice dari awal memang slice adalah referensi tiap elemen
	 */
	var names = [3]string{"ardelia", "pramesti", "cahyaningrum"}
	var firstName = names[0:1] // dari index ke-0 -> index sebelum ke-1
	var allNames = names[:] // semua elemen
	var nameStartFrom1 = names[1:] //semua elemen dimulai dari index ke-1
	var nameEndWith2 = names[:2] //semua elemen sebelum index ke-2
	fmt.Println(firstName, allNames)
	fmt.Println(nameStartFrom1, nameEndWith2)

	firstName[0] = "ardel"
	//index ke-0 dari names akan berubah menjadi ardel
	//karena slice elemen firstName punya memory address yang sama dengan elemen array names
	fmt.Println(names)


	/**
	* Fungsi len(<variabel>)
	* - untuk mengetahui jumlah elemen suatu array/slice
	 */
	var books = []string{"factfulness", "madilog", " gerpolek"}
	fmt.Println(len(books)) // output: 3

	/**
	* Fungsi cap(<variabel>)
	* - untuk mengetahui kapasitas maksimum sebuah array/slice
	 */
	var theNames = []string{"ardi", "daniel", "rully"}
	fmt.Println(len(theNames)) // output: 3
	fmt.Println(cap(theNames)) // output: 3

	var aNames = theNames[0:1]
	fmt.Println(len(aNames)) // output: 1
	fmt.Println(cap(aNames)) // output: 3

	// slicing dimulai dari nilai awal > index ke-0
	// index tersebut akan menjadi index ke-0 baru di slice yang baru
	// contoh names[x:y]
	// jika x > 0 maka x akan menjadi index ke-0 dari slice yang baru
	var bNames = theNames[1:3]
	fmt.Println(len(bNames))
	fmt.Println(cap(bNames))


	fmt.Println("\n=========================")
	/**
	* Fungsi append()
	* - menambahkan elemen apda slice
	 */
	var things = []string{"jam", "televisi", "handphone"}
	var newThings = append(things, "komputer")
	fmt.Println(newThings)

	/**
	* beberapa hal yang perlu diperhatikan
	* - jika len(things) == cap(things), maka elemen baru akan jadi referensi baru
	* - jika len(things) < cap(things), maka elemen baru akan menempati cakupan kapasitas
	 */
	var theThings = things[0:2]
	fmt.Println(len(theThings) < cap(theThings)) // output: true
	var appendThings = append(theThings, "sedotan")
	fmt.Println(things) // output: [jam televisi sedotan]
	fmt.Println(appendThings)

	/**
	* Fungsi copy()
	 */
	dst := make([]string, 3)
	src := []string{"pineapple", "banana", "apple", "watermelon"}
	copyResult := copy(dst, src) // return type int, jumlah yang berhasil di copy
	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(copyResult)

	barang := []string{"jam", "televisi", "kulkas"}
	barangPengcopy := []string{"komputer", "laptop"}
	_ = copy(barang, barangPengcopy)
	// output: [komputer laptop kulkas]
	fmt.Println(barang)
	// jika param 1 dari function copy sudah memiliki elemen, maka
	// elemen tersebut akan direplace

	/**
	* Pengaksesan slice dengan 3 indeks
	* - fungsinya agar bisa menentukan jumlah capacity (cap)
	 */
	var planes = []string{"airbus", "boeing", "bombardier"}
	var aPlanes = planes[0:2] //len: 2, cap: 3
	var bPlanes = planes[0:2:2] //len: 2, cap: 2
	fmt.Println("len", len(aPlanes), "; cap:", cap(aPlanes))
	fmt.Println("len", len(bPlanes), "; cap:", cap(bPlanes))
}

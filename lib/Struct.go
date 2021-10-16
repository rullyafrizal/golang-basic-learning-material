package lib

import "fmt"

func Struct() {
	// Go tidak memiliki class, tetapi memiliki tipe data struktur yang disebut struct
	// Struct adalah kumpulan definisi variabel atau fungsi, yang dibungkus dengan tipe data
	// baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi
	// Mirip seperti Map, hanya saja key-nya sudah didefinisikan di awal,dan tipe data tiap item
	// bisa berbeda

	// Konsep struct mirip dengan konsep OOP di bahasa lain
	// Mirip != sama

	// mengisi value struct
	var student1 student
	student1.name = "Rully Afrizal"
	student1.grade = 12
	student1.major = "science"

	fmt.Println("name : ", student1.name)
	fmt.Println("grade : ", student1.grade)
	fmt.Println("major : ", student1.major)

	// inisialisasi object struct
	// Example 1
	var s1 = student{}
	s1.name = "Ardelia"
	s1.grade = 12
	s1.major = "art"
	//Example 2
	var s2 = student{"Jonathan", 11, "Science"}
	//Example 3
	var s3 = student{name: "Deandra"}

	fmt.Println("s1 : ", s1.name)
	fmt.Println("s2 : ", s2.name)
	fmt.Println("s3 : ", s3.name)

	// Variabel objek Pointer
	var s4 = student{name: "Robert", grade: 10}
	var s5 *student = &s4
	// s5 bukan variabel asli, tetapi nilai propertynya bisa diakses seperti biasa
	// tanpa perlu di-dereferensi
	fmt.Println("s4, name : ", s4.name)
	fmt.Println("s5, name : ", s5.name)

	s5.name = "Downey"
	// di sini name s4 akan juga berubah karena memiliki pointer yang sama
	fmt.Println("s4, name : ", s4.name)
	fmt.Println("s5, name : ", s5.name)


	// Embedded Struct
	// adalah mekanisme menempelkan sebuah struct sebagai properti struct lain
	// agar lebih mudah dipahami
	var e1 = employee{}
	e1.name = "Rudy"
	e1.age = 22
	e1.position = "Backend Developer"
	e1.department = "IT"

	fmt.Println("\nname : ", e1.name +
			"\nage : ", e1.age,
			"\nposition : ", e1.position,
			"\ndepartment : ", e1.department,
		)
	// Embedded struct dengan nama property yang sama
	var p1 = pet{}
	p1.master.name = "Rully"
	p1.master.age = 19
	p1.name = "cukcu"
	p1.age = 1

	fmt.Println("\nmaster name : ", p1.master.name +
		"\nmaster age : ", p1.master.age,
		"\npet name : ", p1.name,
		"\npet age : ", p1.age,
	)

	// Pengisian nilai embedded struct/sub struct
	var m1 = master{"johnny", 25}
	var p2 = pet{master: m1, name: "Chihuahua", age: 2}
	fmt.Println("\nmaster name : ", p2.master.name +
		"\nmaster age : ", p2.master.age,
		"\npet name : ", p2.name,
		"\npet age : ", p2.age,
	)



	// Anonymous struct atau struct stored in variable
	var author = struct {
		person
		phoneNumber string
	}{}

	author.person = person{"Rolf Dobelli", 55}
	author.phoneNumber = "081234567890"
	fmt.Println("\nName : ", author.person.name +
		"\nAge : ", author.person.age,
		"\nPhone Number : ", author.phoneNumber,
	)



	// Kombinasi slice & struct
	var employees = []employee{
		{person{"Rully", 21}, "IT", "Backend Engineer"},
		{person{"Hanif", 22}, "Finance", "Accountant"},
		{person{"Edgar", 25}, "Operation", "Courier"},
		{person{"Vier", 24}, "Support", "Support Staff"},
	}

	for _, employee := range employees {
		fmt.Println("\nname : ", employee.name +
			"\nage : ", employee.age,
			"\nposition : ", employee.position,
			"\ndepartment : ", employee.department,
		)
	}


	// inisialisasi anonymous struct dengan slice
	var pets = []struct{
		pet
	}{
		{pet: pet{master{"rully", 19}, "chikuwa", 2}},
		{pet: pet{master{"indira", 21}, "chibi", 3}},
		{pet: pet{master{"dean", 20}, "kyubi", 1}},
	}

	for _, pet := range pets {
		fmt.Println("\nmaster name : ", pet.master.name +
			"\nmaster age : ", pet.master.age,
			"\npet name : ", pet.name,
			"\npet age : ", pet.age,
		)
	}
}

// Deklarasi struct
type student struct {
	name string
	grade int
	major string
}

// Embedded struct
type person struct {
	name string
	age int
}

type employee struct {
	// person struct is embedded inside employee struct
	person
	department string
	position string
}

// embedded type dengan property yang sama
type master struct {
	name string
	age int
}

type pet struct {
	master
	name string
	age int
}

package _interface

import (
	"fmt"
	"strings"
)

func EmptyInterface() {
	// Mengingat bahwa interface adalah tipe data, sehingga cara penggunaannya sama seperti tipe data lainnya
	var secret interface{}

	secret = "Ethan Hunt"
	fmt.Println(secret)

	secret = []int{1, 2, 3, 4, 5}
	fmt.Println(secret)

	secret = 12.56
	fmt.Println(secret)

	var user map[string]interface{}
	user = map[string]interface{}{
		"name": "rully",
		"age": 12,
		"hobbies": []string{"Swimming", "Aviation"},
	}
	fmt.Println(user)


	// casting variabel interface kosong
	var number interface{}
	number = 5
	var castNumber = number.(int) * 10
	fmt.Println(number, "multiplied by 10 is : ", castNumber)

	var fruits interface{}

	fruits = []string{"Apple", "Mango", "Banana"}
	var gruits = strings.Join(fruits.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")


	// casting variabel interface ksoong ke objek pointer
	var i1 interface{} = &person{name: "Rully", age: 19} // ambil nilai pointernya (referencing)
	var i1Name = i1.(*person).name // casting dengan mengambil nilai aslinya (Dereferencing)
	fmt.Println(i1Name)


	// kombinasi slice, map, dan interface
	var people = []map[string]interface{}{
		{"name": "Rully", "age": 19},
		{"name": "Ardelia", "age": 20},
		{"name": "Rian", "age": 18},
	}

	for _, p := range people {
		fmt.Println("Name :", p["name"], "\nAge:", p["age"])
	}
}

type person struct {
	name string
	age int
}



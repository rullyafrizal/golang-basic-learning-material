package reflect

import (
	"fmt"
	"reflect"
)

func Reflect() {
	// Reflection adalah teknik untuk inspeksi variabel
	// mengambil informasi dari variabel tersebut atau bahkan memanipulasinya
	// Cakupan informasi : struktur variabel, tipe, nilai pointer, dan lain-lain.

	var number = 23
	var reflectNumberValue = reflect.ValueOf(number)
	fmt.Println(reflectNumberValue) // value asli dari number
	fmt.Println("Tipe variabel :", reflectNumberValue.Type())

	if reflectNumberValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel : ", reflectNumberValue.Int())
	}


	// Pengaksesan value dalam bentuk interface
	var number2 = 25
	var reflectVal = reflect.ValueOf(number2)
	fmt.Println("\nTipe variabel :", reflectVal.Type())
	fmt.Println("Nilai variabel :", reflectVal.Interface())
	// Fungsi interface di atas mengembalikan nilai interface kosong
	// nilai aslinya bisa diakses dengan casting interface kosong
	var nilai = reflectVal.Interface().(int)
	fmt.Println("Nilai Asli :", nilai)


	// Pengaksesan informasi property variabel objek
	var p1 = &person{"Rully Afrizal", 19}
	p1.getPropertyInfo()


	// Pengaksesan informasi method variabel objek
	var p2 = &person{"Bunga Citra Lestari", 35}
	fmt.Println("Name :", p2.Name)

	reflectVal = reflect.ValueOf(p2)
	// akses informasi method
	var method = reflectVal.MethodByName("SetName")
	// panggil method
	method.Call([]reflect.Value{
		reflect.ValueOf("Bunga Citra Edited"),
	})
	fmt.Println("Name (new) :", p2.Name)
}


type person struct {
	Name string
	Age int
}

func (p *person) getPropertyInfo() {
	var reflectVal = reflect.ValueOf(p)

	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	var reflectType = reflectVal.Type()

	for i := 0; i < reflectVal.NumField(); i++ {
		fmt.Println(
				"=============================",
				"\nProperty Name :", reflectType.Field(i).Name,
				"\nTipe data :", reflectType.Field(i).Type,
				"\nNilai :", reflectVal.Field(i).Interface(),
			)
	}
}

func (p *person) SetName(newName string) {
	p.Name = newName
}



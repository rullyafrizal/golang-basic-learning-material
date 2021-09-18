package lib

import "fmt"

func DataTypes() {
	// Decimal number
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal : %f\n", decimalNumber)
	fmt.Printf("bilangan desimal : %.4f\n", decimalNumber)

	// boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// string
	var contohString string = "ini adalah contoh string"
	var message string =
		`ini adalah multiline string
		bisa dipakai multiline
		baris ke-3`

	fmt.Printf("%s", contohString)
	fmt.Printf("%s \n", message)

	// Nil(null) dan zero value
	//var zeroVal string = "" //ini adalah zero value dari string
	//var zeroVal2 bool = false //false adalah zero value dari boolean
	//var zeroVal3 uint8 = 0
	//var zeroVal4 = 0.0
	// zero value != nil
}

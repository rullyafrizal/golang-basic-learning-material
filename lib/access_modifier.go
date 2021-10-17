package lib

import "fmt"

func AccessModifier() {
	// ada 2 jenis access modifier di go
	// 1. Hak akses Exported(public)
	// 2. Hak akses Unexported(private)
}

func SayHello() { // Exported
	fmt.Println("Hello this is exported method from access modifier file")
}

func introduce() { // Unexported
	fmt.Println("Hello if you export this function you will get an error")
}

type person struct { // Unexported struct
	name string // Unexported property
	age int // Unexported property
}

type Employee struct { // Exported struct
	Name string // Exported property
	grade int // Unexported property
}



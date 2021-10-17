package _interface

import "fmt"

type hitung2dimensi interface {
	luas() float64
	keliling() float64
}

type hitung3dimensi interface {
	volume() float64
}

type rumusBangunRuang interface {
	hitung2dimensi
	hitung3dimensi
}

func EmbeddedInterface() {
	var bangunRuang rumusBangunRuang

	// ambil nilai referensinya dengan menggunakan &
	bangunRuang = &kubus{20}
	fmt.Println("luas : ", bangunRuang.luas())
	fmt.Println("Keliling : ", bangunRuang.keliling())
	fmt.Println("Volume : ", bangunRuang.volume())
}

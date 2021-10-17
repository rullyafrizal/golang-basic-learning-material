package _interface

import "math"

type kubus struct {
	sisi float64
}

// menggunakan variabel pointer di tiap-tiap objek struct
func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}





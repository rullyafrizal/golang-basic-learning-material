package lib

import "fmt"

func Loop() {
	/**
	* di golang, perulangan hanya "for"
	* tetapi kemampuannya merupakan gabungan for, foreach, dan while
	 */

	// Perulangan menggunakan keyword "for"
	for i := 1; i <= 5; i++ {
		fmt.Println("angka", i)
	}

	// Perulangan keyword "for", dengan argumen hanya kondisi (mirip while)
	var i = 1
	for i <= 5 {
		fmt.Println("Angka", i)
		i++
	}

	// Perulangan keyword "for", tanpa argumen
	var a = 1
	for {
		fmt.Println("Angka a:", a)
		a++
		if a == 6 {
			break
		}
	}

	//keyword break dan continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			if i < 8 {
				continue
			} else {
				break
			}
		}
		fmt.Println("Angka ini adalah", i)
	}

	//nested loop
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// label dalam looping
outerLoop:
	for i:= 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 4 {
				break outerLoop //ini akan menghentikan looping ber-label outerloop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

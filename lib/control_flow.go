package lib

import "fmt"

func ControlFlow() {
	var point = 2

	// if-else
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("lulus tapi remidi")
	} else {
		fmt.Println("gak lulus bos, wong nilaimu mek", point)
	}

	// variabel temporary
	var nilai = 18840.0
	if percent := nilai / 100; percent >= 100 { //percent di sini sebagai variabel temp
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	//switch-case
	var a = 6
	switch a {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	case 4:
		fmt.Println("empat")
	case 5:
		fmt.Println("lima")
	case 6:
		fmt.Println("enam")
	case 7:
		fmt.Println("tujuh")
	case 8:
		fmt.Println("delapan")
	default:
		fmt.Println("default")
	}

	// case untuk banyak kondisi
	var b = 6

	switch b {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			// kurung kurawal 2 untuk statement/instruction lebih dari 1 line
		}
	}

	// contoh lain
	switch {
	case b == 6:
		fmt.Println("oke ini 6")
	case (b < 8) && (b > 3):
		fmt.Println("oke ini awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//fallthrough
	switch {
	case b != 6:
		fmt.Println("oke ini bukan 6")
	case (b < 8) && (b > 3):
		fmt.Println("oke ini awesome")
		fallthrough
	case b == 6:
		fmt.Println("fallthrough")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//nested control flow
	point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}


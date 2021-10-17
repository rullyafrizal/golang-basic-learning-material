package lib

import "fmt"

func Method() {
	// Method adalah function yang menempel pada type (bisa struct atau lainnya)
	var user1 = User{name: "Rully Afrizal", email: "rully@gmail.com", phoneNumber: "081234567890"}

	user1.sayHello()

	var user2 = User{grades: []float32{90, 95}}

	fmt.Println(user2.avgGrade())


	var s1 = Student{"John Wick", "Science", 79}
	fmt.Println("s1 before", s1.name)

	s1.changeName1("Jason Bourne")
	fmt.Println("s1 after", s1.name)

	s1.changeName2("Ethan Hunt")
	fmt.Println("s1 after", s1.name)


	// mengakses method dari variabel objek pointer
	var u1 = &User{grades: []float32{90, 95, 97, 92}}
	fmt.Println(u1.avgGrade())
}

type User struct {
	name string
	email string
	phoneNumber string
	grades []float32
}

func (u User) sayHello() {
	fmt.Println("Hello ", u.name)
}

func (u User) avgGrade() float32 {
	var total float32
	var length = len(u.grades)

	for _, grade := range u.grades {
		total += grade
	}

	return total / float32(length)
}


// Method pointer
type Student struct {
	name string
	major string
	grade int
}

func (s Student) changeName1(name string) {
	fmt.Println("\n---> on changeName1, name changed to", name)
	s.name = name
}

func (s *Student) changeName2(name string) {
	fmt.Println("\n---> on changeName2, name changed to", name)
	s.name = name
}

package init

import "fmt"

var Student = struct {
	Name string
	Grade int
}{}

func init() {
	Student.Name = "Rully Afrizal"
	Student.Grade = 12

	fmt.Println("---> lib/init/func_init.go imported")
}

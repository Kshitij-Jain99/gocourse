package GO_Basics

import "fmt"

var middleName = "Cane"

// declare firstName at package level if you want to use it in multiple functions
var firstName = "Michael"

func main() {
	fmt.Println(middleName)
	fmt.Println(firstName) // now works
}

func printname() {
	// local variable (this shadows the global one if uncommented)
	// firstName := "Michael"
	fmt.Println(firstName)
}

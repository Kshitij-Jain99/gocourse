package intermediate

import (
	"fmt"
	"math"
)

// Declaration:Eg-1
type geometry interface {
	//2 undefined fxns
	area() float64
	perim() float64
}

// Circle:
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) diameter() float64 { // Additional method (not part of geometry interface)
	return 2 * c.radius
}

// Rectangle:
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

// Rectangle1: Not implementing perim()
type rect1 struct {
	width, height float64
}

func (r rect1) area() float64 {
	return r.height * r.width
}

// func (r rect1) perim() float64 {
// 	return 2 * (r.height + r.width)
// }
//After doing it, no error

func measure(g geometry) { //easy print for both r ankd c
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func interfaces_main() {

	//Eg-1
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	// r1 := rect1{width: 3, height: 4}
	measure(r)
	measure(c)
	// measure(r1)
	/*
			When trying to use rect1 in measure(rect1{}), Go shows an error:
			cannot use rect1 as geometry value in argument to measure:
		    rect1 does not implement geometry (missing perim method)
			Reason: rect1 only implemented one method → does not fully satisfy geometry.
			Once both methods (area + perim) are implemented, rect1 now satisfies geometry.
		    It can now be passed to measure() without error.
	*/

	//Eg-2
	myPrinter(1, "John", 45.9, true)

	printType(9)
	printType("John")
	printType(false)

}

// Eg-2: Empty interfaces
// Type Switching with Interfaces: You can check the underlying type of an interface{} value.
func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Type: Unknown")
	}
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

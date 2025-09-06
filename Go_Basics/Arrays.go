package GO_Basics

import "fmt"

func arrays() {

	//Declaration
	//var arrayName [size]elementType

	//Accesing and Updating elements
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	numbers[0] = 9
	fmt.Println(numbers) // [9 0 0 0 20]

	//Array Initialization with Values: Inline Initialization
	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)     // [Apple Banana Orange Grapes]
	fmt.Println("Third element:", fruits[2]) // Orange

	//Arrays are Value Types(copies, not references). Assignment creates a copy, not a reference.
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray //Assigning an array to a new variable → creates a copy.
	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray) // [1 2 3]
	fmt.Println("Copied array:", copiedArray)     // [100 2 3]

	//Iterating Over Arrays
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	//Range-based loop:
	//Short form: for i, v := range numbers { ... } where i is index and v is value
	// underscore is blank identifier, used to store unused values
	for _, v := range numbers {
		fmt.Printf(" Value: %d\n", v)
	}

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println("Array1 is equal to Array2:", array1 == array2) //Array comparison: Same length, same ele type, same ele at each pos

	//Multi-Dimensional Arrays
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	//Pointers with Arrays (Intro)
	//Arrays are copied by default. To reference original, use pointers:
	origArray := [3]int{1, 2, 3}
	var copyArray *[3]int = &origArray
	(*copyArray)[0] = 100

	fmt.Println("Original array:", origArray)
	// fmt.Println("Copied array:", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}

package intermediate

import (
	"errors"
	"fmt"
)

// Eg-1: Custom Error with errors.New
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: square root of negative number")
	}
	// compute the square root
	return 1, nil // placeholder instead of actual sqrt
}

// Eg-2: Function Returning Only error
func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	// Process data
	return nil
}

// Eg-3: Defining Custom Errors (via Structs)
type myError struct {
	message string
}

func (m *myError) Error() string { // Implement the error interface
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

// Eg-4: Using fmt.Errorf
func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}

func errors_main() {

	//Eg-1: Custom Error with errors.New
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	result1, err1 := sqrt(-16)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(result1)

	//Eg-2: Function Returning Only error
	data := []byte{}
	// if err := process(data); err != nil {    //Short if Error Handling Syntax
	err2 := process(data)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println("Data Processed Successfully")

	//Eg-3:--- error interface of builtin package
	err3 := eprocess()
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	println("")

	//Eg-4: Using fmt.Errorf
	err4 := readData()
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	fmt.Println("Data read successfully.")

}

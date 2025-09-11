package intermediate

import (
	"errors"
	"fmt"
)

func custom_errors_main() {

	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Operation completed successfully!")

}

// CustomError struct with code and message
type customError struct {
	code    int
	message string
	er      error //wrapped error
}

// Error returns the error message. Implementing Error() method of error interface.
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.er)
}

// Function that return a custom error: Before adding wrapped error
//	func doSomething() error {
//		return &customError{
//			code:    500,
//			message: "Something went wrong!",
//		}
//	}

// Example function that returns a custom error
func doSomething() error { // Higher-level function that wraps the error
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err, // wrap the lower-level error
		}
	}
	return nil
}

func doSomethingElse() error { // Simulates a lower-level function
	return errors.New("internal error")
}

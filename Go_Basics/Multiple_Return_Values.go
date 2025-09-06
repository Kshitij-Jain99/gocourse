package GO_Basics

import (
	"errors"
	"fmt"
)

func returnValues() {

	// func functionName(parameter1 type1, parameter2 type2,...) (returnType1, returnType2,...){
	//code block
	// return returvalue1, returnValue2,...
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v. Remainder: %v\n", q, r) //Quotient: 3, Remainder: 1

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (quotient int, remainder int) { //(int, int) if unnamed return values
	quotient = a / b
	remainder = a % b
	return // return quotient, remainder - If unnamed return values
	/*
	   Named Return Values
	     - Instead of just return types, you can name return variables in the signature.
	     - Go will automatically return their values when using a plain return.
	*/
}

/*
Error Handling with Multiple Returns:
  - Convention in Go: functions return (value, error).
  - If an error occurs → return zero value + error.
  - If no error occurs → return value + nil.
*/
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}

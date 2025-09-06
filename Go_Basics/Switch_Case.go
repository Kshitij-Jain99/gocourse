package GO_Basics

import "fmt"

func switchCase() {
	/*
	   	Switch statement in go is (switch case default) (fallthrough)
	   	switch expression {
	   	case value1:
	   	      Code to be executed if expression equals value1
	   	      fallthrough   //can be in any case except default
	   	case value2:
	   	      Code to be executed if expression equals value2
	   	case value3:
	   	      Code to be executed if expression equals value3
	   	default:
	   	      Code to be executed if expression does not match any value
	   	}

	   	Switch statement in other languages (switch case break default)
	   	switch expression {
	   	case value1:
	   	     Code to be executed if expression equals value1
	   	     break;
	   	case value2:
	   	     Code to be executed if expression equals value2
	   	     break;
	   	case value3:
	   	     Code to be executed if expression equals value3
	   	     break;
	   	default:
	   	     Code to be executed if expression does not match any value
	   	     break;
	   	}

	   	In many languages, you must write break after each case.
	       In Go, break is implicit → execution stops after the first match.
	       To continue to just next case, use fallthrough.
	*/

	fruit := "pineapple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple.")
	case "banana":
		fmt.Println("It's a banana.")
	default:
		fmt.Println("Unknown Fruit!")
	}

	//Multiple Values in One Case
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Sunday":
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("Invalid day.")
	}

	//Expressions in Cases
	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	//fallthrough explicitly continues to next case.
	//case written just after fallthrough will be executed and then stop.
	num := 2

	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not Two")
	} //Out.: Greater than 1 , Number is 2 - both will be printed.

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

//Type Switch
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case int32: //both int and int32 are different types
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}

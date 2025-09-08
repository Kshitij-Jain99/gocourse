package intermediate

import "fmt"

func closures() {
	sequence := adder() //adder called once and store result in sequence, sequence now holds the inner function (closure).
	//i is initialized to 0 only once, when adder() is called.
	fmt.Println(sequence()) //1
	fmt.Println(sequence()) //2, Each call to sequence() updates and returns the persistent i value.
	fmt.Println(sequence()) //3
	fmt.Println(sequence()) //4
	fmt.Println(sequence()) //5

	//A new i is created (separate state). This operates independently from the first closure.
	subtractor := func() func(int) int { //Anonymous Function Returning Another Function
		countdown := 99
		return func(x int) int { //Inner function: takes an integer parameter x, returns an integer.
			countdown -= x
			return countdown
		}
	}()
	// Using the closure subtracter
	fmt.Println(subtractor(1)) // 98
	fmt.Println(subtractor(2)) // 96
	fmt.Println(subtractor(3)) // 93
	fmt.Println(subtractor(4)) // 89
	fmt.Println(subtractor(5)) // 84

}

func adder() func() int { //No parameters. Returns another function (closure). The returned function itself has no parameters but returns an integer.
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int { //anonymous function (closure) that captures the variable i from its surrounding scope.
		i++
		fmt.Print("added 1 to i: ")
		return i
	}
}

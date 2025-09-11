package intermediate

import (
	"fmt"
	"math/rand"
)

func randomNumbers_main() {

	// // val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(rand.Intn(6) + 5)
	// // fmt.Println(val.Intn(101))
	// fmt.Println(rand.Float64()) // between 0.0 and 1.0

	//Dice Game:
	/*
			Game Design:
		     - Two dice → values between 1 and 6 (inclusive).
		     - Sum of dice → range 2 to 12 (inclusive).
		     - User chooses: Roll the dice or Exit
	*/
	for {
		// Show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int //Input stored

		//Error Handling
		_, err := fmt.Scan(&choice) //pointer-based input.
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		//Dice Roll Logic:
		die1 := rand.Intn(6) + 1 //ensures dice values between 1 and 6.
		die2 := rand.Intn(6) + 1

		// show the results
		fmt.Printf("You rolled a %d and a %d.\n", die1, die2)
		fmt.Println("Total:", die1+die2)

		// Play again feature: Ask of the user wants to roll again
		fmt.Println("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}

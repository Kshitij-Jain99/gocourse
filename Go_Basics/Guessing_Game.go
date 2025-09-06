package GO_Basics

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Computer generates a random number between 1–100.
User repeatedly guesses until correct.
Hints guide user if guess is too high/low.
Loop continues until correct guess → then break.
*/
func logic() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generate random number between 1-100
	target := random.Intn(100) + 1

	//Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		//Check if guess is correct
		if guess == target {
			fmt.Println("congratulations! You guessed the correct number:")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}
}

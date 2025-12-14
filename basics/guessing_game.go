package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingGame() {
	fmt.Println("Hello from guessing game")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	for {
		var guess int
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You've guessed the correct number:", target)
			break
		}
	}

}

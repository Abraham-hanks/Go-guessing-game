package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessGame(secretNumber int) string {

	var guess int
	fmt.Println("Guess a number")
	fmt.Println("Enter your guessed number:")
	fmt.Scan(&guess)
	fmt.Printf("You guessed %d \n", guess)
	if guess == secretNumber {
		answer := "You guessed correct"
		return answer
	} else if guess < secretNumber {
		answer := "Too low, guess again "
		fmt.Println(answer)
		return guessGame(secretNumber)
	} else {
		answer := "Too high, guess again "
		fmt.Println(answer)
		return guessGame(secretNumber)
	}
}

func main() {
	fmt.Printf("hello, Africa\n")

	//generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)
	fmt.Println(guessGame(secretNumber))
}

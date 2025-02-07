package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("We will show a random number beteween 0 and 100, try to guess")

	scanner := bufio.NewScanner(os.Stdin)
	guessing := [10]int64{}

	x := rand.Int64N(101)

	for i := range guessing {
		fmt.Print("Enter a number: ")
		scanner.Scan()

		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInteger, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("You need to enter a integer number")
			return
		}

		switch {
		case guessInteger < x:
			fmt.Println("Wrong number. The number is greather than ", guessInteger)
		case guessInteger > x:
			fmt.Println("Wrong number. The number is less than ", guessInteger)
		case guessInteger == x:
			fmt.Printf("Congrats your guessing is correct the number %d with %d try`s\n", x, i)
			fmt.Printf("Guesses %v\n", guessing[:i])
			return
		}

		guessing[i] = guessInteger
	}

	fmt.Printf("Unfortunetly you don`t guess the number. The correct number is %d\n", x)
	fmt.Printf("After 10 try`s: %v\n", guessing)
}

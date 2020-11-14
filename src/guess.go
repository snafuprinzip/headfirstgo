// guess is the guessing game of numbers
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false

	// 10 Versuche
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("\nYou have", 10-guesses, "left.")
		guess := -1
		// solange loopen bis man eine sinnvolle eingabe erhaelt
		for {
			fmt.Print("Make a guess: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			input = strings.TrimSpace(input)
			guess, err = strconv.Atoi(input)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Illegal number entered:", err)
				continue
			}
			break
		}

		// Auswertung
		if guess < target {
			fmt.Println("Your guess is too low.")
		} else if guess > target {
			fmt.Println("Your guess is too high.")
		} else {
			success = true
			fmt.Println("Good job, jou've guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess the number", target)
	}
}

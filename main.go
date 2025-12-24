package main

import (
	"fmt"
	"math/rand/v2"
	"number-guesser/internal"
	"os"
)

func main() {
	fmt.Println("--------------------------------------------------------")
	fmt.Println("------------------Number guessing game------------------")
	fmt.Println("--------------------------------------------------------")
	fmt.Println()
	var choice int
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println("Enter your choice")
	fmt.Scanln(&choice)
	for i := 1; true; i++ {
		fmt.Println("---------Round", i, "---------")
		randNumb := rand.IntN(100) + 1
		fmt.Println(randNumb)
		switch {
		case choice == 1:
			internal.NumberGuess(randNumb, 10)
		case choice == 2:
			internal.NumberGuess(randNumb, 5)
		case choice == 3:
			internal.NumberGuess(randNumb, 3)
		}
		var exit string
		fmt.Println("Would you like to exit? (y/n)")
		fmt.Scanln(&exit)
		if exit == "y" {
			os.Exit(0)
		}
	}
}

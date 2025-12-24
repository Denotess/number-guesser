package internal

import "fmt"

func NumberGuess(target int, chances int) {
	for i := 1; i <= chances; i++ {
		var guess int
		fmt.Println("Pick you number")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println("Correct")
			return
		}
		if guess > target {
			fmt.Println("Your guess is greater than the target!")
		}
		if guess < target {
			fmt.Println("Your guess is less than the target!")
		}
	}
	fmt.Println("U lost, u've used all of your chances!")
}

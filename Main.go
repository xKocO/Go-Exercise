package main

import (
	"fmt"

	"github.com/xKocO/Go-Exercise.git/Functions/Dice"
	"github.com/xKocO/Go-Exercise.git/Functions/TimeNow"
)

func main() {
	Program()
}

func Program() {
	var x int
	var called bool = true

	if called {
		fmt.Println("Write 1 if you wanna roll a d20, write 2 if you want to know your local time!")
	}
	fmt.Scanln(x)

	if x == 1 {
		Dice.Rolld20()
	} else if x == 2 {
		TimeNow.GiveTime()
	} else {
		fmt.Println("Please input a correct answer!")
		called = false
		Program()
	}
}

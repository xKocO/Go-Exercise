package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xKocO/Go-Exercise.git/Functions/Dice"
	"github.com/xKocO/Go-Exercise.git/Functions/TimeNow"
)

var called bool = true

func main() {
	Program()
}

func Program() {
	if called {
		fmt.Println("Write 'Dice' if you wanna roll a d20, write 'Time' if you want to know your local date and time!")
	}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	Action(input)
}

func Action(x string) {
	x = strings.TrimSpace(x)
	if x == "Dice" {
		Dice.Rolld20()
	} else if x == "Time" {
		TimeNow.GiveTime()
	} else {
		fmt.Println("Please input a correct answer!")
		called = false
		Program()
	}
}

package main

import (
	"fmt"
	"time"
	//"./Functions/Dice"
)

func main() {
	PrintTime()
	//Dice.Rolld20()
}

func PrintTime() {
	fmt.Println("Hello!", "\n", "The current time in your region is:", time.Now())
}

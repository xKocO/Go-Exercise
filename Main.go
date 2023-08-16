package main

import (
	"fmt"
	"time"
)

func main() {
	PrintTime()
	Rolld20()
}

func PrintTime() {
	fmt.Println("Hello!", "\n", "The current time in your region is:", time.Now())
}

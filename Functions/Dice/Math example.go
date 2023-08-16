package Dice

import (
	"fmt"
	"math/rand"
	"time"
)

func Rolld20() {
	fmt.Println("You rolled a d20: ", rand.Intn(20))
	time.Sleep(4 * time.Second)
}

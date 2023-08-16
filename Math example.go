package Math_example

import (
	"fmt"
	"math/rand"
)

func Rolld20() {
	fmt.Println("You rolled a d20: ", rand.Intn(20))
}

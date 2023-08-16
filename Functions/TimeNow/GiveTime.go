package TimeNow

import (
	"fmt"
	"time"
)

func GiveTime() {
	fmt.Println("The local time is: \n", time.Now())
}

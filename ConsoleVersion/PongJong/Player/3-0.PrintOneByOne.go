package Player

import (
	"fmt"

	. "github.com/littletrainee/MahJong/ConsoleVersion"
)

func Printonebyone(value string, index int, length int, suffix string) {
	PrintRGB(value)
	if index != length-1 {
		fmt.Printf(", ")
	} else {
		fmt.Printf(suffix)
	}
}

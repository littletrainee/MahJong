package Player

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
)

func Printonebyone(value string, index int, length int, suffix string) {
	CV.PrintRGB(value)
	if index != length-1 {
		fmt.Printf(", ")
	} else {
		fmt.Printf(suffix)
	}
}

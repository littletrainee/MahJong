package Player

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

func printonebyone(value string, index int, length int, suffix string) {
	CV.PrintRedTextOrNot(value, CC.Tilemap)
	if index != length-1 {
		fmt.Printf(", ")
	} else {
		fmt.Printf(suffix)
	}
}

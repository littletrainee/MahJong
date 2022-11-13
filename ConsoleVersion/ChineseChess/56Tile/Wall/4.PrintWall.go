package Wall

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

func (w *Wall) PrintWall() {
	count := 0
	last := len(w.Wall.Get())
	fmt.Printf("%s: ", w.Name.Get())
	for _, v := range w.Wall.Get() {
		count++
		CV.PrintRedTextOrNot(v, CC.Tilemap)
		if count != last {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n\n")
}

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
		// if v[1] == 'r' {
		// 	// color.RedString(Tilemap.Get()[v])
		// 	RedString.Printf("%v", CCM32TCVBTP.Tilemap.Get()[v])
		// } else {
		// 	fmt.Printf("%v", CCM32TCVBTP.Tilemap.Get()[v])
		// }
		CV.PrintRedTextOrNot(v, CC.Tilemap)
		if count != last {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n\n")
}

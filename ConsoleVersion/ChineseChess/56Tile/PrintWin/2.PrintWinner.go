package PrintWin

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

func (pw *PrintWin) PrintWinner() {
	// fmt.Println(pw.name.Get())
	for i, v := range pw.hand.Get() {
		CV.PrintRedTextOrNot(v, CC.Tilemap)
		if i != len(pw.hand.Get())-1 {
			fmt.Printf(", ")
		}
	}
	if len(pw.meld.Get()) != 0 {
		fmt.Print("     ")
		for _, meld := range pw.meld.Get() {
			fmt.Printf(" [ ")
			for i, v := range meld {
				CV.PrintRedTextOrNot(v, CC.Tilemap)
				if i != len(meld)-1 {
					fmt.Printf(", ")
				} else {
					fmt.Printf(" ]")
				}
			}
		}
	}

	fmt.Print("    ")
	CV.PrintRedTextOrNot(pw.lastone.Get(), CC.Tilemap)
	fmt.Printf("\n\n")
}

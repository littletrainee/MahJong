package PrintWin

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"
)

func (pw *PrintWin) Print() {
	fmt.Printf("               %s\n\n", pw.name.Get())
	switch len(pw.meld.Get()) {
	case 1:
		fmt.Print("  ")
	case 0:
		fmt.Print("     ")
	}
	for i, v := range pw.hand.Get() {
		Player.Printonebyone(v, i, len(pw.hand.Get()), ", ")
	}
	if len(pw.meld.Get()) != 0 {
		fmt.Print("     ")
		for _, meld := range pw.meld.Get() {
			fmt.Printf(" [ ")
			for i, v := range meld {
				Player.Printonebyone(v, i, len(meld), " ]")
			}
		}
	}

	fmt.Print("       ")
	CV.PrintRGB(pw.lastone.Get())
	fmt.Printf("\n\n")
}

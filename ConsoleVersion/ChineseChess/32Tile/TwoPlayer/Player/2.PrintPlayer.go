package Player

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

// Print Player Hand And River
func (p *Player) PrintPlayer() {
	fmt.Printf("%s's Hand: ", p.Name.Get())
	for i, v := range p.Hand.Get() {
		CV.PrintRedTextOrNot(v, CC.Tilemap)
		if i != len(p.Hand.Get())-1 {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
	fmt.Printf("%s's Meld: ", p.Name.Get())
	for _, meld := range p.Meld.Get() {
		fmt.Printf("[ ")
		for i, v := range meld {
			CV.PrintRedTextOrNot(v, CC.Tilemap)
			if i != len(meld)-1 {
				fmt.Printf(", ")
			} else {
				fmt.Printf(" ]")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n%s's River: ", p.Name.Get())
	for i, v := range p.River.Get() {
		CV.PrintRedTextOrNot(v, CC.Tilemap)
		if i != len(p.River.Get())-1 {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
	fmt.Printf("\n")
}

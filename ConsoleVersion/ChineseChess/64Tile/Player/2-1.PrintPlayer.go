package Player

import (
	"fmt"
)

// Print Player Hand And River
func (p *Player) PrintPlayer() {
	fmt.Printf("%s's Hand: ", p.Name.Get())
	for i, v := range p.Hand.Get() {
		printonebyone(v, i, len(p.Hand.Get()), "\n")
	}
	fmt.Printf("%s's Meld: ", p.Name.Get())
	for _, meld := range p.Meld.Get() {
		fmt.Printf("[ ")
		for i, v := range meld {
			printonebyone(v, i, len(meld), " ]")
		}
	}
	fmt.Println()
	fmt.Printf("%s's River: ", p.Name.Get())
	for i, v := range p.River.Get() {
		printonebyone(v, i, len(p.River.Get()), "\n")
	}
	fmt.Printf("\n\n")
}

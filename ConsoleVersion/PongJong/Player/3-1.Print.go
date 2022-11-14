package Player

import "fmt"

// Print Player hand, meld and river
func (p *Player) Print() {
	fmt.Printf("%s's Hand: ", p.Name.Get())
	for i, v := range p.Hand.Get() {
		Printonebyone(v, i, len(p.Hand.Get()), "\n")
	}
	fmt.Printf("%s's Meld: ", p.Name.Get())
	for _, meld := range p.Meld.Get() {
		fmt.Printf("[ ")
		for i, v := range meld {
			Printonebyone(v, i, len(meld), " ]")
		}
	}
	fmt.Println()
	fmt.Printf("%s's River: ", p.Name.Get())
	for i, v := range p.River.Get() {
		Printonebyone(v, i, len(p.River.Get()), "\n")
	}
	fmt.Printf("\n\n")
}

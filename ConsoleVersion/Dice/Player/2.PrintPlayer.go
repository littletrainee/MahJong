package Player

import "fmt"

// Print Player Hand And River
func (p *Player) PrintPlayer() {
	fmt.Printf("%s's Hand: ", p.Name.Get())
	fmt.Printf("%v\n", p.Hand.Get())
	fmt.Printf("%s's River: ", p.Name.Get())
	fmt.Printf("%v\n\n", p.River.Get())
}

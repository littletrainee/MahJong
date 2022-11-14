package Player

import . "github.com/littletrainee/MahJong/ConsoleVersion/PongJong"

// Sort Player Hand
func (p *Player) SortHand() {
	// declare variable and Get variable from Player Hand
	var newtemp []string
	// check value is in hand, then append to newtemp
	for i := range Tile.Get() {
		for j := range p.Hand.Get() {
			if Tile.Get()[i] == p.Hand.Get()[j] {
				newtemp = append(newtemp, p.Hand.Get()[j])
			}
		}
	}
	p.Hand.Set(newtemp)
}

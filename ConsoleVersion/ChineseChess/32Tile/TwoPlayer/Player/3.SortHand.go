package Player

import CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"

// Sort Player Hand
func (p *Player) SortHand() {
	// declare variable and Get variable from Player Hand
	var newtemp []string
	// check value is in hand, then append to newtemp
	for i := range CC.Tile_key {
		for j := range p.Hand.Get() {
			if CC.Tile_key[i] == p.Hand.Get()[j] {
				newtemp = append(newtemp, p.Hand.Get()[j])
			}
		}
	}
	p.Hand.Set(newtemp)
}

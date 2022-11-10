package Player

import (
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
	"github.com/littletrainee/slices"
)

func (p *Player) MakeSmallKangMeld(target string, w *Wall.Wall) {
	var (
		temphand     []string   = p.Hand.Get()
		tempmeldlist [][]string //= p.Meld.Get()
		wall         []string   = w.Wall.Get()
		index        int
		twokand      uint8 = p.TwoKang.Get()
	)
	for _, meld := range p.Meld.Get() {
		if slices.ContainsElement(meld, target) {
			meld = append(meld, target)

			// Remove target from hand
			index = slices.FindIndexOfElement(temphand, target)
			temphand = append(temphand[:index], temphand[index+1:]...)

			// append element from wall last one to temphand
			temphand = append(temphand, wall[len(wall)-1])

			// Remove last element from wall
			wall = wall[:len(wall)-1]
		}
		tempmeldlist = append(tempmeldlist, meld)
	}

	// twokand ++
	twokand++

	// Set each value
	w.Wall.Set(wall)
	p.Hand.Set(temphand)
	p.Meld.Set(tempmeldlist)
	p.TwoKang.Set(twokand)
}

package Player

import (
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/Wall"
	"github.com/littletrainee/slices"
)

func (p *Player) MakeConcealedKangMeld(target string, w *Wall.Wall) {
	var (
		temphand     []string   = p.Hand.Get()
		tempmeldlist [][]string = p.Meld.Get()
		tempmeld     []string
		wall         []string = w.Wall.Get()
		index        int
		twokand      uint8 = p.TwoKang.Get()
	)
	if key := askMeld(p); key == "k" {

		//Add Four Element to tempmeld
		for i := 0; i < 4; i++ {
			tempmeld = append(tempmeld, target)
		}
		// Remove three element from temphand
		for i := 0; i < 4; i++ {
			index = slices.FindIndexOfElement(temphand, target)
			temphand = append(temphand[:index], temphand[index+1:]...)
		}

		// append element from wall last one to temphand
		temphand = append(temphand, wall[len(wall)-1])

		// Remove last element from wall
		wall = wall[:len(wall)-1]

		// append tempmeld to tempmeldlist
		tempmeldlist = append(tempmeldlist, tempmeld)

		// twokand ++
		twokand++

		// Set each value
		p.HasKangMeld.Set(false)
		w.Wall.Set(wall)
		p.Hand.Set(temphand)
		p.Meld.Set(tempmeldlist)
		p.TwoKang.Set(twokand)
	}
}

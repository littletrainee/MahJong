package Player

import (
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
	"github.com/littletrainee/slices"
)

func (p *Player) MakeBigKangMeld(otherplayer *Player, w *Wall.Wall) {
	var (
		temphand         []string   = p.Hand.Get()
		tempmeldlist     [][]string = p.Meld.Get()
		tempmeld         []string
		otherplayerriver []string = otherplayer.River.Get()
		wall             []string = w.Wall.Get()
		index            int
		twokand          uint8 = p.TwoKang.Get()
	)
	//Add Four Element to tempmeld
	for i := 0; i < 4; i++ {
		tempmeld = append(tempmeld, otherplayerriver[len(otherplayerriver)-1])
	}
	// Remove three element from temphand
	for i := 0; i < 3; i++ {
		index = slices.FindIndexOfElement(temphand, otherplayerriver[len(otherplayerriver)-1])
		temphand = append(temphand[:index], temphand[index+1:]...)
	}
	// append element from wall last one to temphand
	temphand = append(temphand, wall[len(wall)-1])

	// remove last element from otherplayerriver
	otherplayerriver = otherplayerriver[:len(otherplayerriver)-1]

	// append tempmeld to tempmeldlist
	tempmeldlist = append(tempmeldlist, tempmeld)

	// twokand ++
	twokand++

	// Set each value
	p.HasPongMeld.Set(false)
	w.Wall.Set(wall)
	otherplayer.River.Set(otherplayerriver)
	p.Hand.Set(temphand)
	p.Meld.Set(tempmeldlist)
	p.IsConcealed.Set(false)
	p.TwoKang.Set(twokand)
}

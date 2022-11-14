package TileType

import (
	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"
	"github.com/littletrainee/slices"
)

func sortHand(hand *GetSet.Type[[]string]) {
	var p Player.Player
	p.Hand.Set(hand.Get())
	p.SortHand()
	hand.Set(p.Hand.Get())
}

func (tt *TileType) Create(thisplayer, otherplayer *Player.Player) {
	temp := thisplayer.Hand.Get()
	if len(thisplayer.Meld.Get()) != 0 {
		for _, meld := range thisplayer.Meld.Get() {
			temp = append(temp, meld[:3]...)
		}
	}
	if len(temp) != 9 {
		temp = append(temp, otherplayer.River.Get()[len(otherplayer.River.Get())-1])
	}
	temp = slices.RemoveDuplicate(temp)

	if len(temp) == 1 {
		tt.colorlength.Set(1)
		tt.kindlength.Set(1)
	} else if len(temp) == 9 {
		tt.colorlength.Set(3)
		tt.kindlength.Set(9)
	} else {
		color, kind := splitColorAndKind(temp)
		tt.colorlength.Set(color)
		tt.kindlength.Set(kind)

	}
}

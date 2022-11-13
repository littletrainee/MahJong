package TileType

import (
	"strconv"

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
		tt.target.Set(strconv.Itoa(len(temp)) + "Color" +
			strconv.Itoa(len(temp)) + "Kind")
	} else if len(temp) == 9 {
		tt.target.Set(strconv.Itoa(len(temp)/3) + "Color" +
			strconv.Itoa(len(temp)) + "Kind")
	} else {
		color, kind := splitColorAndKind(temp)
		tt.target.Set(strconv.Itoa(color) + "Color" + strconv.Itoa(kind) + "Kind")
	}
}

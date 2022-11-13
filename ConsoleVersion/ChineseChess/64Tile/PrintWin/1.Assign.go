package PrintWin

import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/Player"

func (pt *PrintWin) Assign(thisplayer, otherplayer *Player.Player) {
	pt.name.Set(thisplayer.Name.Get())
	pt.meld.Set(thisplayer.Meld.Get())
	pt.istsumo.Set(thisplayer.IsTsumo.Get())
	if !pt.istsumo.Get() { // Ron
		pt.hand.Set(thisplayer.Hand.Get())
		pt.lastone.Set(otherplayer.River.Get()[len(otherplayer.River.Get())-1])
	} else { // Tsumo
		pt.hand.Set(thisplayer.Hand.Get()[:len(thisplayer.Hand.Get())-1])
		pt.lastone.Set(thisplayer.Hand.Get()[len(thisplayer.Hand.Get())-1])
	}
}

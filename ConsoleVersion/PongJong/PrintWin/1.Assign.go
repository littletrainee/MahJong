package PrintWin

import "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"

func (pw *PrintWin) Assign(thisplayer, otherplayer *Player.Player) {
	pw.name.Set(thisplayer.Name.Get())
	pw.meld.Set(thisplayer.Meld.Get())
	pw.istsumo.Set(thisplayer.IsTsumo.Get())
	if !pw.istsumo.Get() { // Ron
		pw.hand.Set(thisplayer.Hand.Get())
		pw.lastone.Set(otherplayer.River.Get()[len(otherplayer.River.Get())-1])
	} else { // Tsumo
		pw.hand.Set(thisplayer.Hand.Get()[:len(thisplayer.Hand.Get())-1])
		pw.lastone.Set(thisplayer.Hand.Get()[len(thisplayer.Hand.Get())-1])
	}
}
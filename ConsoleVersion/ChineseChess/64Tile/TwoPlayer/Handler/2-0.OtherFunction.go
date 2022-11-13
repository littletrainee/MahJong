package Handler

import (
	"sync"

	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/Player"
)

func checkPongAndKang(otherplayerriver string, thisplayer *Player.Player, wg *sync.WaitGroup) {
	// pong
	thisplayer.CheckPong(otherplayerriver, wg)
	if thisplayer.HasPongMeld.Get() {
		wg.Add(1)
		thisplayer.CheckBigKang(otherplayerriver, wg)
	}
}

// repeat check player can concealed or small kang
func repeatcheckkangwithwin(thisplayer, otherplayer *Player.Player, h *Handler) string {
	for {
		thisplayer.TsumoCheck()     // check thisplayer is win
		if thisplayer.Iswin.Get() { // is win
			h.GameState.GameOn.Set(false)
			thisplayer.IsTsumo.Set(true)
			thisplayer.WinOnTheWallTail.Set(true)
			h.PrintWin.Assign(thisplayer, otherplayer)
			return thisplayer.Name.Get()
		} else if target := thisplayer.CheckConcealedKang(); target != "" { // not win and chcek concealed kang
			thisplayer.MakeConcealedKangMeld(target, &h.Wall)
			continue
		} else if target := thisplayer.CheckSmallKang(); target != "" { // concealed target is empty then check small kang
			thisplayer.MakeSmallKangMeld(target, &h.Wall)
			continue
		} else { // not win and no concealed and no small kang
			break
		}
	}
	return ""
}

func checkisron(thisplayer, otherplayer *Player.Player, h *Handler) string {
	h.GameState.GameOn.Set(false)
	h.PrintWin.Assign(thisplayer, otherplayer)
	h.RongBy = *otherplayer
	return thisplayer.Name.Get()
}

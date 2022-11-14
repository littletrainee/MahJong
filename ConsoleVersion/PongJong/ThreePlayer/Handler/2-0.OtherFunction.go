package Handler

import (
	"sync"

	. "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"
)

func checkisron(thisplayer, otherplayer *Player, h *Handler) string {
	h.GameState.GameOn.Set(false)
	h.PrintWin.Assign(thisplayer, otherplayer)
	h.RongBy = *otherplayer
	return thisplayer.Name.Get()
}

func checkPongAndKang(otherplayerriver string, thisplayer *Player, wg *sync.WaitGroup) {
	wg.Add(1)
	// pong
	thisplayer.CheckHasPongMeld(otherplayerriver, wg)
}

func checkpongbygoroutine(p2 *Player, riverlastelement string, h *Handler) {
	go func() {
		if !p2.IsRiiChi.Get() {
			p2.CheckHasPongMeld(riverlastelement, h.wg)
		} else {
			h.wg.Done()
		}
	}()
}

func getmeldlength(p2 *Player, p1 *Player) int {
	meldlength := len(p2.Meld.Get())
	if p2.AskMakePong() == "p" {
		p2.MakePongMeld(p1)
	}
	p2.HasPongMeld.Set(false)
	return meldlength
}

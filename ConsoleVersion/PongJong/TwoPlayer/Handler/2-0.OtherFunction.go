package Handler

import (
	"sync"

	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"
)

func checkisron(thisplayer, otherplayer *Player.Player, h *Handler) string {
	h.GameState.GameOn.Set(false)
	h.PrintWin.Assign(thisplayer, otherplayer)
	h.RongBy = *otherplayer
	return thisplayer.Name.Get()
}

func checkPongAndKang(otherplayerriver string, thisplayer *Player.Player, wg *sync.WaitGroup) {
	wg.Add(1)
	// pong
	thisplayer.CheckPong(otherplayerriver, wg)
}

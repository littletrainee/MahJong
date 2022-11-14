package Handler

import (
	"fmt"

	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/TileType"
)

func (h *Handler) Update() {
	for h.GameState.GameOn.Get() {
		if len(h.Wall.Wall.Get()) == 0 {
			fmt.Println("This Game Is Draw.")
			h.Winner = "Draw"
			break
		}
		h.Del.Run()
		switch h.GameState.GameTurn.Get() {
		case 0:
			h.Winner = h.update(&h.Player1, &h.Player2)
		case 1:
			h.Winner = h.update(&h.Player2, &h.Player1)
		}
		h.Del.Run()
	}
	switch h.Winner {
	case h.Player1.Name.Get():
		h.PrintWin.Print()
		h.tt.Create(&h.Player1, &h.RongBy)
	case h.Player2.Name.Get():
		h.PrintWin.Print()
		h.tt.Create(&h.Player2, &h.RongBy)
	default:
		h.tt = TileType.TileType{}
	}
	h.tt.SetYakuAndPoint()
	if h.Winner != "Draw" {
		h.tt.Print()
	}
}

package Handler

import (
	"fmt"

	. "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/TileType"
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
			h.Winner = h.update(&h.Player1, &h.Player2, &h.Player3)
		case 1:
			h.Winner = h.update(&h.Player2, &h.Player3, &h.Player1)
		case 2:
			h.Winner = h.update(&h.Player3, &h.Player1, &h.Player2)
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
	case h.Player3.Name.Get():
		h.PrintWin.Print()
		h.tt.Create(&h.Player3, &h.RongBy)
	default:
		fmt.Println("DRAW")
		h.tt = TileType{}
	}
	h.tt.SetYakuAndPoint()
	if h.Winner != "Draw" {
		h.tt.Print()
	}
}

package Handler

import (
	"fmt"

	player "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TileType"
)

func (h *Handler) update(p1, p2, p3, p4 *player.Player) string {
	// Check p1 Is Tsumo
	p1.TsumoCheck()
	if !p1.Iswin.Get() {
		h.GameState.NextRound(p1.Name.Get())
		//if tenpai is true
		if p1.TenPai.Get() {
			p1.AutoDiscard()
		} else {
			if p := p1.TenPaiCheck(); len(p) != 0 { // check p1 can declare tenpai
				p1.AskDeclareTenPai() // ask declare tenpai
			}
			// discard
			p1.Discard()
			// Sort p1 Hand after discard
			p1.SortHand()
		}
		h.Del.Run()
		h.wg.Add(3)
		go p2.RonCheck(p1, h.wg)
		go p3.RonCheck(p1, h.wg)
		go p4.RonCheck(p1, h.wg)
		h.wg.Wait()
		riverlastelement := p1.River.Get()[len(p1.River.Get())-1]
		// p1 Discard To Let p4 Ron
		if p4.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			h.PrintWin.Assign(p4, p1)
			return p4.Name.Get()
		} else if p3.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			h.PrintWin.Assign(p3, p1)
			return p3.Name.Get()
		} else if p2.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			h.PrintWin.Assign(p2, p1)
			return p2.Name.Get()
		} else {
			p2.Iswin.Set(false)
			p3.Iswin.Set(false)
			p4.Iswin.Set(false)
			p2.IsTsumo.Set(false)
			p3.IsTsumo.Set(false)
			p4.IsTsumo.Set(false)
			if !p2.TenPai.Get() {
				p2.CheckChi(riverlastelement) // check has chi meld
				if p2.HasMeld.Get() != nil {
					choice := p2.AskMakeMeldAndChoose() // ask player yes or not to make meld
					if choice != 0 {                    // want make meld and made decision
						p2.MakeMeld(choice, p1)
						goto TurnToNextPlayer
					} else {
						goto DrawCard
					}
				}
			}
		DrawCard:
			// And Draw card
			p2.DrawCard(&h.Wall)
			// check wall is the last one
			if len(h.Wall.Wall.Get()) == 0 {
				h.Wall.LastTile.Set(true)
			}
		TurnToNextPlayer:
			// Set GameTurn to p2
			h.GameState.TurnNext()
		}
	} else {
		h.GameState.GameOn.Set(false)
		p1.IsTsumo.Set(true)
		h.PrintWin.Assign(p1, p4)
		return p1.Name.Get()
	}
	return ""
}

func (h *Handler) Update() {
	for h.GameState.GameOn.Get() {
		if len(h.Wall.Wall.Get()) == 0 {
			fmt.Println("This Game Is Draw.")
			h.Winner = "Draw"
			break
		}
		h.Del.Run()
		switch h.GameState.GameTurn.Get() {
		case 0: // Player1
			h.Winner = h.update(&h.Player1, &h.Player2, &h.Player3, &h.Player4)
		case 1: // Player2
			h.Winner = h.update(&h.Player2, &h.Player3, &h.Player4, &h.Player1)
		case 2:
			h.Winner = h.update(&h.Player3, &h.Player4, &h.Player1, &h.Player2)
		case 3:
			h.Winner = h.update(&h.Player4, &h.Player1, &h.Player2, &h.Player3)
		}
	}
	h.PrintWin.PrintWinner()

	switch h.Winner {
	case h.Player1.Name.Get():
		h.tt.Create(&h.Player1, &h.Player4, &h.Wall, &h.GameState)
	case h.Player2.Name.Get():
		h.tt.Create(&h.Player2, &h.Player1, &h.Wall, &h.GameState)
	case h.Player3.Name.Get():
		h.tt.Create(&h.Player3, &h.Player2, &h.Wall, &h.GameState)
	case h.Player4.Name.Get():
		h.tt.Create(&h.Player4, &h.Player3, &h.Wall, &h.GameState)
	default:
		h.tt = TileType.TileType{}
	}
	if h.Winner != "Draw" {
		h.tt.PrintTileType()
	}
}

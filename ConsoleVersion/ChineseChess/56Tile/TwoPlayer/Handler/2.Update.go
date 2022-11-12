package Handler

import (
	"fmt"
	"sync"

	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/TileType"
)

func checkPongAndKang(otherplayerriver string, thisplayer *Player.Player, wg *sync.WaitGroup) {
	// pong
	wg.Add(1)
	go thisplayer.CheckPong(otherplayerriver, wg)
	wg.Wait()
	if thisplayer.HasPongMeld.Get() {
		wg.Add(1)
		go thisplayer.CheckBigKang(otherplayerriver, wg)
		wg.Wait()
	}
}

func (h *Handler) update(p1, p2 *Player.Player) string {
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
		h.wg.Add(int(h.GameState.Maxplayer.Get()) - 1)
		go p2.RonCheck(p1, h.wg)
		h.wg.Wait()
		riverlastelement := p1.River.Get()[len(p1.River.Get())-1]
		// p1 Discard To Let p2 Ron
		if p2.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			h.PrintWin.Assign(p2, p1)
			return p2.Name.Get()
		} else {
			p2.Iswin.Set(false)
			p2.IsTsumo.Set(false)
			if !p2.TenPai.Get() {
				// check has pong meld
				checkPongAndKang(riverlastelement, p2, h.wg)
				// check has chi meld
				p2.CheckChi(riverlastelement)
				choice := p2.AskMakeMeldAndChoose() // ask player yes or not to make meld
				switch choice {
				case "c1", "c2", "c3":
					p2.MakeChiMeld(choice, p1)
					goto TurnToNextPlayer
				case "p":
					p2.MakePongMeld(p1)
					goto TurnToNextPlayer
				case "k":
					p2.MakeBigKangMeld(p1, &h.Wall)
					goto TurnToNextPlayer
				case "s":
					goto Drawcard
				}
			}
		Drawcard:
			p2.DrawCard(&h.Wall)
			// check wall is the last one
			if len(h.Wall.Wall.Get()) == 0 {
				h.Wall.LastTile.Set(true)
			}
		TurnToNextPlayer:
			// check player can concealed or small kang
			for {
				// p2.TsumoCheck()
				if p2.Iswin.Get() {
					h.GameState.GameOn.Set(false)
					p2.IsTsumo.Set(true)
					p2.WinOnTheWallTail.Set(true)
					h.PrintWin.Assign(p2, p1)
					return p2.Name.Get()
				}
				if target := p2.CheckConcealedKang(); target != "" {
					p2.MakeConcealedKangMeld(target, &h.Wall)
					goto TurnToNextPlayer
				}
				if target := p2.CheckSmallKang(); target != "" {
					p2.MakeSmallKangMeld(target, &h.Wall)
				} else {
					break
				}
			}
			// Set GameTurn to p2
			h.GameState.TurnNext()
		}
	} else {
		h.GameState.GameOn.Set(false)
		p1.IsTsumo.Set(true)
		h.PrintWin.Assign(p1, p2)
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
			h.Winner = h.update(&h.Player1, &h.Player2)
		case 1: // Player2
			h.Winner = h.update(&h.Player2, &h.Player1)
		}
	}

	switch h.Winner {
	case h.Player1.Name.Get():
		h.PrintWin.PrintWinner()
		h.tt.Create(&h.Player1, &h.Player2, &h.Wall, &h.GameState)
	case h.Player2.Name.Get():
		h.PrintWin.PrintWinner()
		h.tt.Create(&h.Player2, &h.Player1, &h.Wall, &h.GameState)
	default:
		h.tt = TileType.TileType{}
	}
	if h.Winner != "Draw" {
		h.tt.PrintTileType()
	}
}

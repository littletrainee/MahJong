package Handler

import "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/Player"

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
		// p1 Discard To Let p2, p3 and p4 Ron
		h.wg.Add(int(h.GameState.Maxplayer.Get()) - 1)
		go p2.RonCheck(p1, h.wg)
		h.wg.Wait()
		riverlastelement := p1.River.Get()[len(p1.River.Get())-1]
		if p2.Iswin.Get() {
			return checkisron(p2, p1, h)
		} else {
			// check p4, p3 and p2 has pong or kang meld
			if !p2.TenPai.Get() {
				h.wg.Add(1)
				go checkPongAndKang(riverlastelement, p2, h.wg)
			}
			h.wg.Wait()
			if !p2.TenPai.Get() {
				// p2 check has chi meld
				p2.CheckChi(riverlastelement)
				choice := p2.AskMakeMeldAndChoose() // ask player yes or not to make meld
				switch choice {
				case "c1", "c2", "c3":
					p2.MakeChiMeld(choice, p1)
					p2.ResetMeld()
					goto TurnToNextPlayer
				case "p":
					p2.MakePongMeld(p1)
					p2.ResetMeld()
					goto TurnToNextPlayer
				case "k":
					p2.MakeBigKangMeld(p1, &h.Wall)
					if p2iswin := repeatcheckkangwithwin(p2, p1, h); p2iswin == "" {
						p2.ResetMeld()
						goto TurnToNextPlayer
					} else {
						return p2iswin
					}
				case "s":
					p2.ResetMeld()
					goto Drawcard
				}
			}
		Drawcard:
			p2.DrawCard(&h.Wall)
			// check wall is the last one
			if len(h.Wall.Wall.Get()) == 0 {
				h.Wall.LastTile.Set(true)
			}
			if p2iswin := repeatcheckkangwithwin(p2, p1, h); p2iswin == "" { // check p2 is win
				p2.ResetMeld()
				goto TurnToNextPlayer // not win then go to TurnToNextPlayer label
			} else {
				return p2iswin
			}
		TurnToNextPlayer:
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

package Handler

import "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"

func (h *Handler) update(p1, p2 *Player.Player) string {
	// Check p1 Is Tsumo
	p1.TsumoCheck()
	if !p1.Iswin.Get() {
		h.GameState.NextRound(p1.Name.Get())
		if p1.IsRiiChi.Get() {
			p1.AutoDiscard()
		} else {
			if p := p1.RiiChiCheck(); len(p) != 0 { // check p1 can declare riichi
				p1.AskDeclareRiiChi() // ask declare riichi
			}
			p1.ManualDiscard() // manual discard
			p1.SortHand()      // sort p1 Hand after manual discard
		}
		h.Del.Run()
		h.wg.Add(int(h.GameState.Maxplayer.Get()) - 1)
		go p2.RonCheck(p1, h.wg)
		h.wg.Wait()
		riverlastelement := p1.River.Get()[len(p1.River.Get())-1]
		if p2.Iswin.Get() {
			return checkisron(p2, p1, h)
		} else {
			h.wg.Add(1)
			checkpongbygoroutine(p2, riverlastelement, h)
			h.wg.Wait()
			meldlength := getmeldlength(p2, p1)
			if meldlength != len(p2.Meld.Get()) {
				h.GameState.TurnTargetPlayer(p2.CodeNumber.Get())
			} else {
				p2.DrawCard(&h.Wall)
				h.GameState.TurnNext()
			}
		}
	} else {
		h.GameState.GameOn.Set(false)
		p1.IsTsumo.Set(true)
		h.PrintWin.Assign(p1, p2)
		return p1.Name.Get()
	}
	return ""
}

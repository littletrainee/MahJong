package Handler

import . "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"

func (h *Handler) update(p1, p2, p3 *Player) string {
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
		go p3.RonCheck(p1, h.wg)
		h.wg.Wait()
		riverlastone := p1.River.Get()[len(p1.River.Get())-1]
		if p3.Iswin.Get() {
			return checkisron(p3, p1, h)
		} else if p2.Iswin.Get() {
			return checkisron(p2, p1, h)
		} else {
			h.wg.Add(int(h.GameState.Maxplayer.Get()) - 1)
			checkpongbygoroutine(p3, riverlastone, h)
			checkpongbygoroutine(p2, riverlastone, h)
			h.wg.Wait()
			if p2.HasPongMeld.Get() {
				meldlength := getmeldlength(p2, p1)
				if meldlength != len(p2.Meld.Get()) { // has make meld
					h.GameState.TurnTargetPlayer(p2.CodeNumber.Get())
					goto ReturnEmptyString
				}
			}
			if p3.HasPongMeld.Get() {
				p3meldlength := getmeldlength(p3, p1)
				if p3meldlength != len(p3.Meld.Get()) { // has make meld
					h.GameState.TurnTargetPlayer(p3.CodeNumber.Get())
					goto ReturnEmptyString
				}
			}
			p2.DrawCard(&h.Wall)
			h.GameState.TurnNext()
		}
	} else {
		h.GameState.GameOn.Set(false)
		p1.IsTsumo.Set(true)
		h.PrintWin.Assign(p1, p2)
		return p1.Name.Get()
	}
ReturnEmptyString:
	return ""
}

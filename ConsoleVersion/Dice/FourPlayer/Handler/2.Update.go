package Handler

import (
	"fmt"

	player "github.com/littletrainee/MahJong/ConsoleVersion/Dice/FourPlayer/Player"
)

func (h *Handler) update(p1, p2, p3, p4 *player.Player) {
	// Check p1 Is Tsumo
	p1.TsumoCheck()
	if !p1.Iswin.Get() {
		h.GameState.NextRound(p1.Name.Get())
		p1.Discard()
		p1.SortHand()
		h.Del.Run()
		h.wg.Add(3)
		go p2.RonCheck(p1, h.wg)
		go p3.RonCheck(p1, h.wg)
		go p4.RonCheck(p1, h.wg)
		h.wg.Wait()
		if p4.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			fmt.Printf("\n%s is Ron.\nThe Hand is %v.	Win By %v. ",
				p4.Name.Get(), p4.Hand.Get(), p1.River.Get())
		} else if p3.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			fmt.Printf("\n%s is Ron.\nThe Hand is %v.	Win By %v. ",
				p3.Name.Get(), p3.Hand.Get(), p1.River.Get())
		} else if p2.Iswin.Get() {
			h.GameState.GameOn.Set(false)
			fmt.Printf("\n%s is Ron.\nThe Hand is %v.	Win By %v. ",
				p2.Name.Get(), p2.Hand.Get(), p1.River.Get())
		} else {
			p2.Iswin.Set(false)
			p3.Iswin.Set(false)
			p4.Iswin.Set(false)
			// Clear p2 Hand
			p2.Hand.Set([]uint8{})
			// And Redraw card
			for i := 0; i < 5; i++ {
				p2.DrawCard()
			}
			p2.SortHand()
			h.GameState.TurnNext()
		}
	} else {
		h.GameState.GameOn.Set(false)
		fmt.Printf("%s is Tsumo.\nThe Hand is %v.	Win By %v. ",
			p1.Name.Get(), p1.Hand.Get(), p1.Hand.Get()[len(p1.Hand.Get())-1])
	}
}

func (h *Handler) Update() {
	for h.GameState.GameOn.Get() {
		h.Del.Run()
		switch h.GameState.GameTurn.Get() {
		case 0: // Player1
			h.update(&h.Player1, &h.Player2, &h.Player3, &h.Player4)
		case 1: // Player2
			h.update(&h.Player2, &h.Player3, &h.Player4, &h.Player1)
		case 2: //Player3
			h.update(&h.Player3, &h.Player4, &h.Player1, &h.Player2)
		case 3: // Player4
			h.update(&h.Player4, &h.Player1, &h.Player2, &h.Player3)

		}
	}
}

package Handler

import (
	"math/rand"
	"sync"
	"time"

	CC "github.com/littletrainee/ClearConsole"
)

func (h *Handler) Start() {
	rand.Seed(time.Now().UnixNano())
	h.GameState.GameOn.Set(true)
	h.GameState.GameTurn.Set(0) // Set to Player 1
	h.GameState.SetMaxPlayer(4)
	h.wg = &sync.WaitGroup{}
	// Set Player Name
	h.Player1.Name.Set("Player 1")
	h.Player2.Name.Set("Player 2")
	h.Player3.Name.Set("Player 3")
	h.Player4.Name.Set("Player 4")
	// Each Player Draw Card 2 tile 2 time (4 tile)
	for i := 0; i < 4; i++ {
		h.Player1.DrawCard()
		h.Player2.DrawCard()
		h.Player3.DrawCard()
		h.Player4.DrawCard()
	}
	// Player1 Draw the 5th card
	h.Player1.DrawCard()
	// // Sort Hand
	h.Player1.SortHand()
	h.Player2.SortHand()
	h.Player3.SortHand()
	h.Player4.SortHand()
	// ADD Printfunction to Delegate slice
	h.Del.Add(CC.ClearConsole)
	h.Del.Add(h.Player1.PrintPlayer)
	h.Del.Add(h.Player2.PrintPlayer)
	h.Del.Add(h.Player3.PrintPlayer)
	h.Del.Add(h.Player4.PrintPlayer)
}

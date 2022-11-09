package Handler

import (
	"sync"

	CC "github.com/littletrainee/ClearConsole"
	wall "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Wall"
)

func (h *Handler) Start() {
	wall.Init()
	h.GameState.GameOn.Set(true)
	h.GameState.GameTurn.Set(0) // Set to Player 1
	h.GameState.GameRound.Set(1)
	h.GameState.SetMaxPlayer(2)
	h.wg = &sync.WaitGroup{}
	// Set Player Name
	h.Wall.Name.Set("Wall")
	h.Player1.Name.Set("Player 1")
	h.Player2.Name.Set("Player 2")
	h.Player3.Name.Set("Player 3")
	// combine two slice to one Tilemap
	wall.CreateMap()
	// Add Tile map to Wall Hand
	h.Wall.AppendTileMapToWall()
	// Shuffle Wall Hand
	h.Wall.Shuffle()
	// Each Player Draw Card 2 tile 2 time (4 tile)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			h.Player1.DrawCard(&h.Wall)
		}
		for j := 0; j < 2; j++ {
			h.Player2.DrawCard(&h.Wall)
		}
		for j := 0; j < 2; j++ {
			h.Player3.DrawCard(&h.Wall)
		}
	}
	// Player1 Draw the 5th card
	h.Player1.DrawCard(&h.Wall)
	// // Sort Hand
	h.Player1.SortHand()
	h.Player2.SortHand()
	h.Player3.SortHand()
	// ADD Printfunction to Delegate slice
	h.Del.Add(CC.ClearConsole)
	h.Del.Add(h.GameState.PrintRound)
	h.Del.Add(h.Wall.PrintWall)
	h.Del.Add(h.Player1.PrintPlayer)
	h.Del.Add(h.Player2.PrintPlayer)
	h.Del.Add(h.Player3.PrintPlayer)
}

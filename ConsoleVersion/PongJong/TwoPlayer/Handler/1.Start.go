package Handler

import (
	"sync"

	CC "github.com/littletrainee/ClearConsole"
	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Wall"
)

func (h *Handler) Start() {
	// Initialization Wall Random Seed
	Wall.Init()
	// Set GameState
	h.GameState.GameOn.Set(true) // loop switch
	h.GameState.GameTurn.Set(0)  // start with player 1
	h.GameState.GameRound.Set(0) // Start from 0 round
	h.GameState.SetMaxPlayer(2)  // this time have two player join
	h.wg = &sync.WaitGroup{}

	// Set Name
	h.Wall.Name.Set("Wall")
	h.Player1.Name.Set("Player 1")
	h.Player2.Name.Set("Player 2")
	// Set Player Code Number
	h.Player1.CodeNumber.Set(1)
	h.Player2.CodeNumber.Set(2)
	// Set Player 1 is bookmaker
	h.Player1.Bookmaker.Set(true)
	// Combine Tile_Color and Tile_Kind to Tile
	h.Wall.CreateTile()
	// Append Tile to Wall
	h.Wall.AppendTileToWall()
	// Shuffle Wall
	h.Wall.Shuffle()
	// Each Player Draw Card 2 Tile 4 time (8 Tile)
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			h.Player1.DrawCard(&h.Wall)
		}
		for j := 0; j < 2; j++ {
			h.Player2.DrawCard(&h.Wall)
		}
	}
	// Player 1  Draw the 9th tile
	h.Player1.DrawCard(&h.Wall)
	// Sort Hand
	h.Player1.SortHand()
	h.Player2.SortHand()
	// Add PrintFunction to Delegate slice
	h.Del.Add(CC.ClearConsole)
	h.Del.Add(h.GameState.PrintRound)
	h.Del.Add(h.Wall.Print)
	h.Del.Add(h.Player1.Print)
	h.Del.Add(h.Player2.Print)
}

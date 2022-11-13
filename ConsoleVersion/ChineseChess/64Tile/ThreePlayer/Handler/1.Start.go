package Handler

import (
	"sync"

	CC "github.com/littletrainee/ClearConsole"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/Wall"
)

func (h *Handler) Start() {
	Wall.Init()
	h.GameState.GameOn.Set(true)
	h.GameState.GameTurn.Set(0)
	h.GameState.GameRound.Set(1)
	h.GameState.SetMaxPlayer(3)
	h.wg = &sync.WaitGroup{}
	// Set Player Name
	h.Wall.Name.Set("Wall")
	h.Player1.Name.Set("Player 1")
	h.Player2.Name.Set("Player 2")
	h.Player3.Name.Set("Player 3")
	// Set Player Code Number
	h.Player1.CodeNumber.Set(1)
	h.Player2.CodeNumber.Set(2)
	h.Player3.CodeNumber.Set(3)
	h.Player1.IsConcealed.Set(true)
	h.Player2.IsConcealed.Set(true)
	h.Player3.IsConcealed.Set(true)
	h.Player1.Bookmaker.Set(true)
	h.Player1.ContinueToBookmaker.Set(0)
	// combine tile_key and tile_value slice to one tilemap
	Wall.CreateMap()
	// Add Tile map to Wall Wall
	h.Wall.AppendTileMapToWall()
	// Suffle Wall
	h.Wall.Shuffle()
	// Each Player Draw Card 2 Tile 3 time (6 Tile)
	for i := 0; i < 3; i++ {
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
	// playerdraw the 7th tile
	h.Player1.DrawCard(&h.Wall)
	h.Player2.DrawCard(&h.Wall)
	h.Player3.DrawCard(&h.Wall)
	// player1 draw the 8th card
	h.Player1.DrawCard(&h.Wall)
	// Sort Hand
	h.Player1.SortHand()
	h.Player2.SortHand()
	h.Player3.SortHand()
	// Add Printfunction to Delegate slice
	h.Del.Add(CC.ClearConsole)
	h.Del.Add(h.GameState.PrintRound)
	h.Del.Add(h.Wall.PrintWall)
	h.Del.Add(h.Player1.PrintPlayer)
	h.Del.Add(h.Player2.PrintPlayer)
	h.Del.Add(h.Player3.PrintPlayer)
}

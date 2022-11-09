package Handler

import (
	"sync"

	"github.com/littletrainee/Delegate"
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	player "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Player"
	TT "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TileType"
	wall "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Wall"
)

type Handler struct {
	Wall                      wall.Wall
	Player1, Player2, Player3 player.Player
	Del                       Delegate.ParameterNoneAndReturnNone
	GameState                 CV.GameState
	wg                        *sync.WaitGroup
	Winner                    string
	tt                        TT.TileType
}

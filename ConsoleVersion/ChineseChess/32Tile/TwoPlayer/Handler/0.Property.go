package Handler

import (
	"sync"

	"github.com/littletrainee/Delegate"
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	player "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TwoPlayer/Player"
	TT "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TwoPlayer/TileType"
	wall "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TwoPlayer/Wall"
)

type Handler struct {
	Wall             wall.Wall
	Player1, Player2 player.Player
	Del              Delegate.ParameterNoneAndReturnNone
	GameState        CV.GameState
	wg               *sync.WaitGroup
	Winner           string
	tt               TT.TileType
}

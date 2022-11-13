package Handler

import (
	"sync"

	"github.com/littletrainee/Delegate"
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/PrintWin"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/TileType"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
	"github.com/littletrainee/MahJong/Interface"
)

type Handler struct {
	Wall             Wall.Wall
	Player1, Player2 Player.Player
	Del              Delegate.ParameterNoneAndReturnNone
	GameState        CV.GameState
	wg               *sync.WaitGroup
	Winner           string
	tt               TileType.TileType
	PrintWin         PrintWin.PrintWin
	Interface.IHandler
}

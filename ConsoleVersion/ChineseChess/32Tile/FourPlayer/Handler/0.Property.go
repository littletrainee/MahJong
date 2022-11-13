package Handler

import (
	"sync"

	"github.com/littletrainee/Delegate"
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	player "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Player"
	PW "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/PrintWin"
	TT "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TileType"
	wall "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Wall"
	"github.com/littletrainee/MahJong/Interface"
)

type Handler struct {
	Wall                               wall.Wall
	Player1, Player2, Player3, Player4 player.Player
	Del                                Delegate.ParameterNoneAndReturnNone
	GameState                          CV.GameState
	wg                                 *sync.WaitGroup
	Winner                             string
	tt                                 TT.TileType
	PrintWin                           PW.PrintWin
	Interface.IHandler
}

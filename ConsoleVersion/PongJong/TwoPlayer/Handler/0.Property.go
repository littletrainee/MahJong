package Handler

import (
	"sync"

	"github.com/littletrainee/Delegate"
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/PrintWin"
	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/TileType"
	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Wall"
	"github.com/littletrainee/MahJong/Interface"
)

type Handler struct {
	Wall                     Wall.Wall
	Player1, Player2, RongBy Player.Player
	Del                      Delegate.ParameterNoneAndReturnNone
	GameState                CV.GameState
	wg                       *sync.WaitGroup
	Winner                   string
	tt                       TileType.TileType
	PrintWin                 PrintWin.PrintWin
	Interface.IHandler
}

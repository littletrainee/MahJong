package Handler

import (
	"sync"

	. "github.com/littletrainee/Delegate"
	. "github.com/littletrainee/MahJong/ConsoleVersion"
	. "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"
	. "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/PrintWin"
	. "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/TileType"
	. "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Wall"
	. "github.com/littletrainee/MahJong/Interface"
)

type Handler struct {
	Wall                              Wall
	Player1, Player2, Player3, RongBy Player
	Del                               ParameterNoneAndReturnNone
	GameState                         GameState
	wg                                *sync.WaitGroup
	Winner                            string
	tt                                TileType
	PrintWin                          PrintWin
	IHandler
}

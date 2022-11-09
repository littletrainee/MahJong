package Handler

import (
	"sync"

	"github.com/littletrainee/Delegate"
	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	player "github.com/littletrainee/MahJong/ConsoleVersion/Dice/Player"
)

type Handler struct {
	Player1, Player2, Player3, Player4 player.Player
	Del                                Delegate.ParameterNoneAndReturnNone
	GameState                          CV.GameState
	wg                                 *sync.WaitGroup
}

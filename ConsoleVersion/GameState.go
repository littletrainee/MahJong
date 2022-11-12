package ConsoleVersion

import (
	"fmt"

	"github.com/littletrainee/GetSet"
)

type GameState struct {
	GameOn    GetSet.Type[bool]
	GameTurn  GetSet.Type[uint8]
	GameRound GetSet.Type[uint8]
	Maxplayer GetSet.Type[uint8]
}

func (gs *GameState) SetMaxPlayer(c int) {
	gs.Maxplayer.Set(uint8(c))
}

func (gs *GameState) TurnNext() {
	temp := gs.GameTurn.Get()
	if temp < gs.Maxplayer.Get()-1 {
		temp++
	} else {
		temp = 0
	}
	gs.GameTurn.Set(temp)
}

func (gs *GameState) TurnTargetPlayer(target uint8) {
	gs.GameTurn.Set(target - 1)
}

func (gs *GameState) NextRound(name string) {
	if name == "Player 1" {
		i := gs.GameRound.Get()
		i++
		gs.GameRound.Set(i)
	}
}

func (gs *GameState) PrintRound() {
	fmt.Printf("第 %d 巡\n", gs.GameRound.Get())
}

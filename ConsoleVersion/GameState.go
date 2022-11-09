package ConsoleVersion

import (
	"fmt"

	"github.com/littletrainee/GetSet"
)

type GameState struct {
	GameOn    GetSet.Type[bool]
	GameTurn  GetSet.Type[uint8]
	GameRound GetSet.Type[uint8]
	maxplayer GetSet.Type[uint8]
}

func (g *GameState) SetMaxPlayer(c int) {
	g.maxplayer.Set(uint8(c))
}

func (g *GameState) TurnNext() {
	temp := g.GameTurn.Get()
	if temp < g.maxplayer.Get() {
		temp++
	} else {
		temp = 0
	}
	g.GameTurn.Set(temp)
}

func (g *GameState) NextRound(name string) {
	if name == "Player 1" {
		i := g.GameRound.Get()
		i++
		g.GameRound.Set(i)
	}
}

func (g *GameState) PrintRound() {
	fmt.Printf("第 %d 巡\n", g.GameRound.Get())
}

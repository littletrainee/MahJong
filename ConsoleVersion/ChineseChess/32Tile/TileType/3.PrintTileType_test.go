package TileType

import (
	"testing"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Wall"
)

func TestPrintTileType(t *testing.T) {
	var (
		T      TileType
		p1, p2 Player.Player
		w      Wall.Wall
		gs     CV.GameState
	)
	p1.Hand.Set([]string{"7r", "7r", "7r", "1b", "1r"})
	p1.IsTsumo.Set(true)
	T.Create(&p1, &p2, &w, &gs)
	T.PrintTileType()
}

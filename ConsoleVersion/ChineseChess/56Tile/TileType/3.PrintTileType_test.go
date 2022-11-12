package TileType

import (
	"testing"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
)

func TestPrintTileType(t *testing.T) {
	var (
		T      TileType
		p1, p2 Player.Player
		w      Wall.Wall
		gs     CV.GameState
	)
	p1.Hand.Set([]string{"1r", "1r", "2r", "2b", "3r", "3r", "4r", "4r"})
	p1.IsTsumo.Set(true)
	T.Create(&p1, &p2, &w, &gs)
	T.PrintTileType()
}

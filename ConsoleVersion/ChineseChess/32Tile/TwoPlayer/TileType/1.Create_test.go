package TileType

import (
	"testing"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TwoPlayer/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/TwoPlayer/Wall"
)

func TestCreatehand(t *testing.T) {
	var (
		k      TileType
		p1, p2 Player.Player
		gs     CV.GameState
		wall   Wall.Wall
	)
	p1.Hand.Set([]string{"1b", "2b", "3b", "4b", "4b"})
	k.Create(&p1, &p2, &wall, &gs)

}

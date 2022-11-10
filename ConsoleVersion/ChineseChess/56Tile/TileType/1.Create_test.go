package TileType

import (
	"testing"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Player"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
)

func TestCreatehand(t *testing.T) {
	var (
		k      TileType
		p1, p2 Player.Player
		gs     CV.GameState
		wall   Wall.Wall
	)
	p1.Hand.Set([]string{"1b", "2b", "3b", "4b", "5b", "6b", "7b"})
	p2.River.Set([]string{"7b"})
	p1.IsConcealed.Set(true)
	k.Create(&p1, &p2, &wall, &gs)

}

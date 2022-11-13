package Player

import (
	"testing"

	wall "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/32Tile/Wall"
)

func TestPrintPlayer(t *testing.T) {
	var p Player
	// var w wall.Wall
	wall.CreateMap()
	p.Name.Set("P")
	p.Hand.Set([]string{"2r", "2r"})
	p.Meld.Set([][]string{{"1b", "2b", "3b"}})
	p.River.Set([]string{"3r", "4r"})
	p.PrintPlayer()
}

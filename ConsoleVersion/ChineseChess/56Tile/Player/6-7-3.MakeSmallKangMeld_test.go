package Player

import (
	"testing"

	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
)

func TestMakeSmallKangMeld(t *testing.T) {
	var (
		p Player
		w Wall.Wall
	)
	p.Hand.Set([]string{"3b", "1b"})
	target := "1b"
	w.Wall.Set([]string{"4b"})
	p.Meld.Set([][]string{{"1b", "1b", "1b"}, {"2b", "2b", "2b"}})
	p.MakeSmallKangMeld(target, &w)
	if len(p.Meld.Get()[0]) != 4 {
		t.Error("Wrong")
	}
}

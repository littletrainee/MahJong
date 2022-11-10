package Player

import (
	"testing"

	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Wall"
)

func TestMakeBigKangMeld(t *testing.T) {
	var (
		p1, p2 Player
		w      Wall.Wall
	)
	p1.Hand.Set([]string{"1b", "1b", "1b", "2b", "2b", "2b", "3b"})
	p2.River.Set([]string{"1b"})
	w.Wall.Set([]string{"2b"})
	p1.MakeBigKangMeld(&p2, &w)
	if len(p1.Meld.Get()[0]) != 4 {
		t.Error("Wrong")
	}
}

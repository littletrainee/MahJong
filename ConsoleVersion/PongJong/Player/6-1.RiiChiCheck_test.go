package Player

import (
	"fmt"
	"testing"

	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Wall"
)

func TestTenPaiCheck(t *testing.T) {
	var (
		p Player
		w Wall.Wall
	)
	w.CreateTile()
	p.Hand.Set([]string{"rb", "rb", "rb", "gb", "gb", "gb", "bb", "bb", "bp"})
	v := p.RiiChiCheck()
	if len(v) == 0 {
		t.Error("Error")
	} else {
		fmt.Println(v)
	}
}

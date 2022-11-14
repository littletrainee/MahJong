package Player

import (
	"testing"

	. "github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Wall"
)

func TestTsumoCheck(t *testing.T) {
	var (
		p Player
		w Wall
	)
	w.CreateTile()
	p.Hand.Set([]string{"rb", "rb", "rb", "gb", "gb", "gb"})
	p.Meld.Set([][]string{{"bb", "bb", "bb"}})
	p.TsumoCheck()
	if p.Iswin.Get() != true {
		t.Error("Not Win")
	}
}

package Player

import "testing"

func TestTsumoCheck(t *testing.T) {
	var p Player
	p.Hand.Set([]string{"1b", "2b", "3b", "1r", "1r"})
	p.Meld.Set([][]string{{"7r", "7r", "7r"}})
	p.TsumoCheck()
	if p.Iswin.Get() != true {
		t.Error("Not Win")
	}
}

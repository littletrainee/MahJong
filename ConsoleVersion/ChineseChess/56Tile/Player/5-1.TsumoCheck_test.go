package Player

import "testing"

func TestTsumoCheck(t *testing.T) {
	var p Player
	p.Hand.Set([]string{"1b", "1b", "2b", "2b", "4r"})
	p.Meld.Set([][]string{{"3r", "3r", "4r"}})
	p.TsumoCheck()
	if p.Iswin.Get() != true {
		t.Error("Not Win")
	}
}

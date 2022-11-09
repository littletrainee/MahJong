package Player

import "testing"

func TestTsumoCheck(t *testing.T) {
	var p Player
	p.Hand.Set([]string{"7b", "7b"})
	p.Meld.Set([][]string{{"7b", "7b", "7b"}})
	p.TsumoCheck()
	if p.Iswin.Get() != true {
		t.Error("Not Win")
	}
}

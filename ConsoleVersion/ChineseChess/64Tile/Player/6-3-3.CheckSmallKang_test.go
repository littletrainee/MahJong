package Player

import "testing"

func TestCheckSmallKang(t *testing.T) {
	var (
		p Player
	)
	p.Hand.Set([]string{"1b", "2b", "2b", "2b", "3b"})
	p.Meld.Set([][]string{{"1b", "1b", "1b"}})
	if p.CheckSmallKang() == "" {
		t.Error("Wrong")
	}
}

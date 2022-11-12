package Player

import "testing"

func TestPrintprobablymeld(t *testing.T) {
	var p Player
	p.HasChiMeld.Set([]string{"1b3b"})
	i := p.AskMakeMeldAndChoose()
	if i != "p" {
		t.Errorf("GotErrors %s not one", i)
	}
}

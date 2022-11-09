package Player

import "testing"

func TestPrintprobablymeld(t *testing.T) {
	var p Player
	p.HasMeld.Set([]string{"1b3b"})
	i := p.AskMakeMeldAndChoose()
	if i != 1 {
		t.Errorf("GotErrors %d not one", i)
	}
}

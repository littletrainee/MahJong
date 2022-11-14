package Player

import "testing"

func TestPrintprobablymeld(t *testing.T) {
	var p Player
	i := p.AskMakePong()
	if i != "p" {
		t.Errorf("GotErrors %s not one", i)
	}
}

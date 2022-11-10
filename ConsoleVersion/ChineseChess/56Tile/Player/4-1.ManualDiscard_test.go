package Player

import "testing"

func TestDiscard(t *testing.T) {
	var p Player
	p.Hand.Set([]string{"1b", "2b"}) //, "3b", "4b", "7b"})
	p.Discard()
	if len(p.Hand.Get()) != 4 || len(p.Hand.Get()) != 1 {
		t.Error("Function Wrong!")
	}

}

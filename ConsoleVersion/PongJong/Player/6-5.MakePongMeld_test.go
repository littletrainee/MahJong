package Player

import "testing"

func TestMakePongMeld(t *testing.T) {
	var (
		p1, p2 Player
	)
	p1.Hand.Set([]string{"1b", "1b", "2b", "2b", "2b", "3b"})
	p1.HasPongMeld.Set(true)
	p2.River.Set([]string{"1b"})
	p1.MakePongMeld(&p2)
}

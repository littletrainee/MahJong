package Player

import "testing"

func TestMakeMeld(t *testing.T) {
	var p Player
	var p2 Player
	p.Hand.Set([]string{"1b", "3b", "4b", "5b"})
	p.HasChiMeld.Set([]string{"1b3b"})
	p2.River.Set([]string{"2b"})
	p.MakeChiMeld("c1", &p2)

}

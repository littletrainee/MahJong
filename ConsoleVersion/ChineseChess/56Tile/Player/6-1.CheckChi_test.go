package Player

import (
	"testing"
)

func TestCheckChi(t *testing.T) {
	var p Player
	hand := []string{"2b", "3b", "4b", "5b"}
	p.Hand.Set(hand)
	p.CheckChi("1b")
}

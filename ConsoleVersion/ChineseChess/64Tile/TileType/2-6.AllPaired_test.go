package TileType

import "testing"

func TestAllPaired(t *testing.T) {
	var k TileType
	k.hand.Set([]string{"3b", "3b"})
	k.meld.Set([][]string{{"1b", "1b", "1b"}, {"2b", "2b", "2b"}})
	k.eye.Set("3b")
	k.AllPaired()
}

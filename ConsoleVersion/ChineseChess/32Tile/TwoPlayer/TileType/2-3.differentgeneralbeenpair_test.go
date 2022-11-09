package TileType

import "testing"

func TestDifferentGeneralBeenPair(t *testing.T) {
	var k TileType
	k.hand.Set([]string{"1b", "2b", "3b", "4b", "1r"})
	k.differentgeneralbeenpair()
}

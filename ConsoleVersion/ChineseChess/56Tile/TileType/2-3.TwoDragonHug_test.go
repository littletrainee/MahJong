package TileType

import "testing"

func TestDifferentGeneralBeenPair(t *testing.T) {
	var k TileType
	k.hand.Set([]string{"1b", "1b", "2b", "2b", "3b", "3b", "4b", "4b"})
	k.eye.Set("4b")
	k.isconcealed.Set(true)
	k.TwoDragonHug()
}

package TileType

import "testing"

func TestSamekind(t *testing.T) {
	var k TileType
	k.hand.Set([]string{"7r", "7r", "7r", "7b", "7b"})
	k.samekind()
}

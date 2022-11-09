package TileType

import "testing"

func TestFivesoldier(t *testing.T) {
	var k TileType
	k.hand.Set([]string{"7r", "7r", "7r", "7r", "7r"})
	k.fivesoldier()
}

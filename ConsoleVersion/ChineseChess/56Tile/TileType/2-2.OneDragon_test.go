package TileType

import "testing"

func TestOneDragon(t *testing.T) {
	var k TileType
	k.hand.Set([]string{"4b", "5b", "6b", "7b", "7b"})
	k.meld.Set([][]string{{"1b", "2b", "3b"}})
	k.OneDragon()
}

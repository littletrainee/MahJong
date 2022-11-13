package TileType

import "testing"

func TestFiveSoldier(t *testing.T) {
	var tt TileType
	tt.hand.Set([]string{"7b", "7b", "7b", "7b", "7b", "1r", "2r", "3r"})
	tt.eye.Set("7b")
	tt.FiveSoldier()

}

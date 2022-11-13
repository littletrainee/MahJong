package TileType

import "testing"

func TestOnlyOrNoGeneralAndSorider(t *testing.T) {
	var tt TileType
	tt.hand.Set([]string{"2b", "2b", "2b", "3b", "3b", "3b", "2r", "2r"})
	tt.OnlyOrNoGeneralAndSorider()

}

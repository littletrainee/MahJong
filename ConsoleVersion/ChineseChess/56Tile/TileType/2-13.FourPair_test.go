package TileType

import "testing"

func TestFourPair(t *testing.T) {
	var tt TileType
	tt.hand.Set([]string{"1b", "1b", "2b", "2b", "3b", "3b", "4b", "4b"})
	tt.isconcealed.Set(true)
	tt.FourPair()
	if tt.total.Get() != 1 {
		t.Error("Wrong")
	}
}

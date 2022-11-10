package TileType

import "testing"

func TestContinuToBookmaker(t *testing.T) {
	var tt TileType
	tt.continuetobookmaker.Set(2)
	tt.bookmaker.Set(true)
	tt.ContinueToBookmaker()
	if tt.total.Get() != 4 {
		t.Error("Wrong")
	}
}

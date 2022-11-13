package TileType

import (
	"fmt"
	"testing"
)

func TestOneDragon(t *testing.T) {
	var tt TileType
	tt.hand.Set([]string{"4b", "5b", "6b", "7b", "7b"})
	tt.meld.Set([][]string{{"1b", "2b", "3b"}})
	tt.OneDragon()
	if tt.total.Get() != 4 {
		t.Error("Wrong")
	} else {
		fmt.Println(tt.total.Get())
	}
}

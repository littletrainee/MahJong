package TileType

import "fmt"

// "天胡"
func (t *TileType) tenhou() {
	if t.gameround.Get() == 1 && t.gameturn.Get() == 0 {
		v := t.total.Get()
		v += 6
		t.total.Set(v)
		fmt.Println("天胡        6台")
	}
}

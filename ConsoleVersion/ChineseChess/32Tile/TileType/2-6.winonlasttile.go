package TileType

import "fmt"

func (t *TileType) winonlasttile() {
	if t.lasttile.Get() {
		v := t.total.Get()
		v += 1
		t.total.Set(v)
		fmt.Println("海底撈月    1台")
	}
}

package TileType

import "fmt"

func (t *TileType) WinOnLastTile() {
	if t.lasttile.Get() {
		v := t.total.Get()
		if t.istsumo.Get() {
			v += 3
			fmt.Println("海底撈月    3台")
		} else {
			v += 1
			fmt.Println("河底撈魚    1台")
		}
		t.total.Set(v)
	}
}

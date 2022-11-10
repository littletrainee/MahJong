package TileType

import "fmt"

func (tt *TileType) WinOnTheWallTail() {
	if tt.winonthewalltail.Get() {
		v := tt.total.Get()
		v += 2
		tt.total.Set(v)
		fmt.Println("槓上開花    2台")
	}
}

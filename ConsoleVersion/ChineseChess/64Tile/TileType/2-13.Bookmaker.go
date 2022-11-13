package TileType

import "fmt"

func (tt *TileType) Bookmaker() {
	if tt.bookmaker.Get() {
		v := tt.total.Get()
		v++
		tt.total.Set(v)
		fmt.Println("莊家        1台")
	}
}

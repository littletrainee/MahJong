package TileType

import "fmt"

// "天胡"
func (tt *TileType) TenHou() {
	if tt.gameround.Get() == 1 && tt.gameturn.Get() == 0 {
		v := tt.total.Get()
		v += 8
		tt.total.Set(v)
		fmt.Println("天胡        8台")
	}
}

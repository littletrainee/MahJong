package TileType

import "fmt"

func (tt *TileType) AllOrHalfRequire() {
	if len(tt.meld.Get()) == 2 && !tt.istsumo.Get() {
		v := tt.total.Get()
		v += 2
		tt.total.Set(v)
		fmt.Println("全求        2台")
	}
	if len(tt.meld.Get()) == 2 && tt.istsumo.Get() {
		v := tt.total.Get()
		v += 1
		tt.total.Set(v)
		fmt.Println("半求        1台")
	}
}

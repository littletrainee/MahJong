package TileType

import "fmt"

func (tt *TileType) TsumoOrRon() {
	v := tt.total.Get()
	if !tt.isconcealed.Get() && tt.istsumo.Get() {
		v += 2
		fmt.Println("自摸        2台")
	} else if !tt.isconcealed.Get() && !tt.istsumo.Get() {
		v++
		fmt.Println("胡牌        1台")
	}
	tt.total.Set(v)
}

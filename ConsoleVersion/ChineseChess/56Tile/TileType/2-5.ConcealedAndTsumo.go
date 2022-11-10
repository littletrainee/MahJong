package TileType

import "fmt"

func (tt *TileType) ConcealedAndTsumo() {
	v := tt.total.Get()
	if tt.isconcealed.Get() && tt.istsumo.Get() {
		v += 3
		fmt.Println("門清自摸    3台")
	} else if tt.isconcealed.Get() && !tt.istsumo.Get() {
		v++
		fmt.Println("門清        1台")
		v++
		fmt.Println("胡牌        1台")
	} else if !tt.isconcealed.Get() && tt.istsumo.Get() {
		v++
		fmt.Println("自摸        1台")
	} else if !tt.isconcealed.Get() && !tt.istsumo.Get() {
		v++
		fmt.Println("胡牌        1台")
	}
	tt.total.Set(v)
}

package TileType

import "fmt"

func (tt *TileType) TwoKang() {
	if tt.twokang.Get() == 2 {
		v := tt.total.Get()
		v += 2
		tt.total.Set(v)
		fmt.Println("二槓子      2台")
	}
}

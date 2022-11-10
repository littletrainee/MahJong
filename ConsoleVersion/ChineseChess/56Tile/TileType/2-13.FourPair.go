package TileType

import "fmt"

func (tt *TileType) FourPair() {
	temp := tt.hand.Get()
	if tt.isconcealed.Get() {
		for i := 0; i < 8; i += 2 {
			if temp[i] != temp[i+1] {
				break
			}
			if i == 6 {
				v := tt.total.Get()
				v++
				tt.total.Set(v)
				fmt.Println("四對子      1台")
			}
		}
	}
}

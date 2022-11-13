package TileType

import "fmt"

func (tt *TileType) EightGenerals() {
	for i, v := range tt.hand.Get() {
		if v != "7b" && v != "7r" {
			break
		}
		if i == len(tt.hand.Get())-1 {
			v := tt.total.Get()
			v += 4
			tt.total.Set(v)
			tt.iseightgenerals.Set(true)
			fmt.Println("八家將      4台")
		}
	}
}

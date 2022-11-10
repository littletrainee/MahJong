package TileType

import (
	"fmt"

	"github.com/littletrainee/GetSet"
)

func (tt *TileType) OnlyOrNoGeneralAndSorider() {
	var (
		temp GetSet.Type[[]string]
	)
	temp.Set(tt.hand.Get())
	appendfrommeld(&temp, &tt.meld)
	// only general and sorider
	for i, v := range temp.Get() {
		if v[0] != '1' && v[0] != '7' {
			break
		}
		if i == len(temp.Get())-1 {
			v := tt.total.Get()
			v += 2
			tt.total.Set(v)
			fmt.Println("將帥領兵    2台")
		}
	}
	for i, v := range temp.Get() {
		if v[0] != '1' && v[0] != '7' {
			if i == len(temp.Get())-1 {
				v := tt.total.Get()
				v += 1
				tt.total.Set(v)
				fmt.Println("斷頭尾      1台")
			} else {
				continue
			}
		} else {
			break
		}
	}
}

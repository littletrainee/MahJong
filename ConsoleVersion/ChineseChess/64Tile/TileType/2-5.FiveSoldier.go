package TileType

import "fmt"

func (tt *TileType) FiveSoldier() {
	temp := tt.hand.Get()
	for _, target := range [][]string{{"7b", "五卒連橫    2台"}, {"7r", "五兵合縱    2台"}} {
		if isfivesoldier(temp, target[0]) {
			v := tt.total.Get()
			v += 2
			tt.total.Set(v)
			fmt.Println(target[1])
			break
		}
	}
}

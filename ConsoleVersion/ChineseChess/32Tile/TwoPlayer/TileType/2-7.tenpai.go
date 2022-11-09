package TileType

import "fmt"

func (t *TileType) tenpai() {
	if t.istenpai.Get() {
		v := t.total.Get()
		v += 1
		t.total.Set(v)
		fmt.Println("聽牌        1台")
	}
}

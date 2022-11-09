package TileType

import "fmt"

func (t *TileType) tsumoorron() {
	if t.istsumo.Get() {
		v := t.total.Get()
		v += 2
		t.total.Set(v)
		fmt.Println("自摸        2台")
	} else {
		v := t.total.Get()
		v += 1
		t.total.Set(v)
		fmt.Println("胡牌        1台")
	}
}

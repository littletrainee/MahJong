package TileType

import (
	"fmt"

	"github.com/littletrainee/GetSet"
)

func tilekind(temp []string, targetkind rune) bool {
	for _, v := range temp {
		if rune(v[1]) != targetkind {
			return false
		}
	}
	return true
}
func (t *TileType) SameKind() {
	var temp GetSet.Type[[]string]
	temp.Set(t.hand.Get())
	appendfrommeld(&temp, &t.meld)
	if tilekind(temp.Get(), 'b') || tilekind(temp.Get(), 'r') {
		v := t.total.Get()
		v += 1
		t.total.Set(v)
		fmt.Println("清一色      1台")
	}
}

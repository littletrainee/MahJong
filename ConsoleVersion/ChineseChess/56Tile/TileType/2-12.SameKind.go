package TileType

import "fmt"

func tilekind(temp []string, targetkind rune) bool {
	for _, v := range temp {
		if rune(v[1]) != targetkind {
			return false
		}
	}
	return true
}
func (t *TileType) SameKind() {
	temp := t.hand.Get()
	if tilekind(temp, 'b') || tilekind(temp, 'r') {
		v := t.total.Get()
		v += 1
		t.total.Set(v)
		fmt.Println("清一色      1台")
	}
}

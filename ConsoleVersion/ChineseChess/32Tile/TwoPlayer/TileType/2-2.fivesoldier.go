package TileType

import "fmt"

func five(temp []string, target string) bool {
	for _, v := range temp {
		if v != target {
			return false
		}
	}
	return true
}

func (t *TileType) fivesoldier() {
	temp := t.hand.Get()
	if five(temp, "7r") || five(temp, "7b") {
		v := t.total.Get()
		v += 5
		t.total.Set(v)
		fmt.Println("五卒連橫    5台")
	}
}

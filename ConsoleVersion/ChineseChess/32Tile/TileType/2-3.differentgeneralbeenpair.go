package TileType

import (
	"fmt"

	"github.com/littletrainee/slices"
)

func (t *TileType) differentgeneralbeenpair() {
	temp := t.hand.Get()
	if slices.ContainsElement("1b", temp) && slices.ContainsElement("1r", temp) {
		v := t.total.Get()
		v += 2
		t.total.Set(v)
		fmt.Println("將帥聽令    2台")
	}
}

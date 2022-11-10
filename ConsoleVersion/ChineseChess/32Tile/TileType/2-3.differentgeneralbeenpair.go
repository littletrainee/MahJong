package TileType

import (
	"fmt"

	"github.com/littletrainee/slices"
)

func (t *TileType) differentgeneralbeenpair() {
	temp := t.hand.Get()
	if slices.ContainsElement(temp, "1b") && slices.ContainsElement(temp, "1r") {
		v := t.total.Get()
		v += 2
		t.total.Set(v)
		fmt.Println("將帥聽令    2台")
	}
}

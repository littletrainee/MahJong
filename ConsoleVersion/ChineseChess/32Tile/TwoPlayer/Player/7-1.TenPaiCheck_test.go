package Player

import (
	"fmt"
	"testing"
)

func TestTenPaiCheck(t *testing.T) {
	var p Player
	p.Hand.Set([]string{"1b", "2b", "3b", "4b", "5b"})
	v := p.TenPaiCheck()
	if len(v) == 0 {
		t.Error("Error")
	} else {
		fmt.Println(v)
	}
}

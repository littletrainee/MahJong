package TileType

import (
	"fmt"
	"testing"

	"github.com/littletrainee/MahJong/ConsoleVersion/PongJong/Player"
)

func TestCreatehand(t *testing.T) {
	var (
		tt     TileType
		p1, p2 Player.Player
	)
	p1.Hand.Set([]string{"rc", "rp", "rb", "bc", "bp", "bb", "gc", "gp"})
	p2.River.Set([]string{"gb"})
	tt.Create(&p1, &p2)
	if tt.yaku.Get() != "3Color9Kind" {
		t.Error("Wrong")
	} else {
		fmt.Println(tt.yaku.Get())
	}

}

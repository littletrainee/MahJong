package TileType

import (
	// "fmt"

	"fmt"

	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Player"
)

func (tt *TileType) AllPaired() {
	var (
		temp GetSet.Type[[]string]
	)
	temp.Set(tt.hand.Get())
	appendfrommeld(&temp, &tt.meld)
	sortHand(&temp)
	temp.Set(Player.RemoveEye(tt.eye.Get(), temp.Get()))
	if Player.Ismeld(temp.Get()[:3]) == "triple" && Player.Ismeld(temp.Get()[3:]) == "triple" {
		v := tt.total.Get()
		v += 1
		tt.total.Set(v)
		fmt.Println("碰碰胡      2台")
	}
}

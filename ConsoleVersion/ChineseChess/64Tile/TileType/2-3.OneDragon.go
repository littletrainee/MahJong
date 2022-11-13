package TileType

import (
	"fmt"

	"github.com/littletrainee/GetSet"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

func (tt *TileType) OneDragon() {
	// declare temphand
	var temphand GetSet.Type[[]string]

	temphand.Set(tt.hand.Get())
	appendfrommeld(&temphand, &tt.meld)
	sortHand(&temphand)

	for _, target := range [][]string{
		append(CC.Tile_key[6:], CC.Tile_key[13]),
		append(CC.Tile_key[:7], CC.Tile_key[6])} {
		if isonedragon(temphand.Get(), target) {
			v := tt.total.Get()
			v += 4
			tt.total.Set(v)
			fmt.Println("一條龍      4台")
			break
		}
	}
}

package TileType

import (
	"fmt"

	"github.com/littletrainee/GetSet"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

func (tt *TileType) OneDragon() {
	// declare temphand
	var (
		temphand GetSet.Type[[]string]
		red      []string
		black    []string
	)
	temphand.Set(tt.hand.Get())
	appendfrommeld(&temphand, &tt.meld)
	sortHand(&temphand)
	red = CC.Tile_key[7:]
	red = append(red, CC.Tile_key[13])
	black = CC.Tile_key[:7]
	black = append(black, CC.Tile_key[6])
	for i := range temphand.Get() {
		if temphand.Get()[i] != red[i] {
			break
		} else if i == len(temphand.Get())-1 {
			v := tt.total.Get()
			v += 4
			tt.total.Set(v)
			fmt.Println("一條龍      4台")

		}
	}
	for i := range temphand.Get() {
		if temphand.Get()[i] != black[i] {
			break
		} else if i == len(temphand.Get())-1 {
			v := tt.total.Get()
			v += 4
			tt.total.Set(v)
			fmt.Println("一條龍      4台")

		}
	}
}

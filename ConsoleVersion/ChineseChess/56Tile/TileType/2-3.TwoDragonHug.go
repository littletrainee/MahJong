package TileType

import (
	"fmt"

	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/56Tile/Player"
)

func (t *TileType) TwoDragonHug() {
	var (
		temp  []string = t.hand.Get()
		meld1 []string
		meld2 []string
	)
	if t.isconcealed.Get() {
		temp = Player.RemoveEye(t.eye.Get(), temp)
		for i := 0; i < 6; i += 2 {
			meld1 = append(meld1, temp[i])
			meld2 = append(meld2, temp[i+1])
		}
		for i := range meld1 {
			if meld1[i] != meld2[i] {
				break
			} else if i == len(meld1)-1 {
				v := t.total.Get()
				v += 4
				t.total.Set(v)
				fmt.Println("雙龍抱      4台")
			}
		}
	}

}

package TileType

import "github.com/littletrainee/GetSet"

func appendfrommeld(hand *GetSet.Type[[]string], meld *GetSet.Type[[][]string]) {
	temp := hand.Get()
	if meld.Get() != nil {
		for _, v := range meld.Get() {
			temp = append(temp, v...)
		}
		hand.Set(temp)
	}
}

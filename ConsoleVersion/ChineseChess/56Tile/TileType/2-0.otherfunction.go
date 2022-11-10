package TileType

import "github.com/littletrainee/GetSet"

func appendfrommeld(hand *GetSet.Type[[]string], meld *GetSet.Type[[][]string]) {
	if meld.Get() != nil {
		for _, v := range meld.Get() {
			hand.Set(append(hand.Get(), v...))
		}
	}
}

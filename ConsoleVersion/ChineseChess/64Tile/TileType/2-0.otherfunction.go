package TileType

import (
	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess/64Tile/Player"
	"github.com/littletrainee/slices"
)

func appendfrommeld(hand *GetSet.Type[[]string], meld *GetSet.Type[[][]string]) {
	temp := hand.Get()
	if meld.Get() != nil {
		for _, v := range meld.Get() {
			temp = append(temp, v...)
		}
		hand.Set(temp)
	}
}

func isonedragon(temphand, targettype []string) bool {
	for i := range temphand {
		if temphand[i] != targettype[i] {
			break
		} else if i == len(temphand)-1 {
			return true
		}
	}
	return false
}

func isfivesoldier(temp []string, eye string) bool {
	if slices.ContainsElement(temp, eye) && slices.CountNumber(temp, eye) == 5 {
		for i := 0; i < 5; i++ {
			index := slices.FindIndexOfElement(temp, eye)
			temp = append(temp[:index], temp[index+1:]...)
		}
		if Player.Ismeld(temp) == "triple" || Player.Ismeld(temp) == "sequence" {
			return true
		}
	}
	return false
}

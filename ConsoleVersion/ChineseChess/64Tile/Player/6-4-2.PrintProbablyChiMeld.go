package Player

import (
	"fmt"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

// print probably meld
func printprobablychimeld(probablymeld []string) {
	fmt.Println()
	for i := 0; i < len(probablymeld); i++ {
		// print index
		fmt.Printf("%d.", i+1)
		CV.PrintRedTextOrNot(probablymeld[i][:2], CC.Tilemap)
		CV.PrintRedTextOrNot(probablymeld[i][2:], CC.Tilemap)
		if i != len(probablymeld)-1 {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
}

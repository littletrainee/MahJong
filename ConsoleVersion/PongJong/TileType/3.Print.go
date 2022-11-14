package TileType

import "fmt"

func (tt *TileType) Print() {
	fmt.Println("===========================================")
	fmt.Printf("役名                             %s\n", tt.yaku.Get())
	fmt.Println("-------------------------------------------")
	fmt.Printf("親                                    %d\n", tt.score.Get()[0])
	fmt.Printf("子                                    %d\n", tt.score.Get()[1])
	fmt.Println("===========================================")
}

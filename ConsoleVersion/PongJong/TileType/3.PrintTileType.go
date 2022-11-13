package TileType

import "fmt"

func (tt *TileType) PrintTileType() {
	fmt.Println("役名        ")
	fmt.Println("---------------")
	fmt.Println(tt.target.Get())

	fmt.Println("===============")
}

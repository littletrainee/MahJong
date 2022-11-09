package TileType

import "fmt"

func (t *TileType) PrintTileType() {
	fmt.Println("牌型        台數")
	fmt.Println("---------------")
	t.tenhou()
	t.winonlasttile()
	t.tsumoorron()
	t.samekind()
	t.differentgeneralbeenpair()
	t.fivesoldier()
	fmt.Println("===============")
	fmt.Printf("共：        %d台\n", t.total.Get())
}

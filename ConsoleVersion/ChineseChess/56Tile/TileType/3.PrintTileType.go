package TileType

import "fmt"

func (tt *TileType) PrintTileType() {
	fmt.Println("牌型        台數")
	fmt.Println("---------------")
	tt.TenHou()
	tt.TenPai()
	tt.Bookmaker()
	tt.ContinueToBookmaker()
	tt.OnlyOrNoGeneralAndSorider()
	tt.ConcealedAndTsumo()
	tt.AllOrHalfRequire()
	tt.WinOnLastTile()
	tt.SameKind()
	if tt.eye.Get() != "4pair" {
		tt.AllPaired()
		tt.TwoKang()
		tt.TwoDragonHug()
		tt.OneDragon()
	} else {
		tt.FourPair()
	}
	fmt.Println("===============")
	fmt.Printf("共：        %d台\n", tt.total.Get())
}

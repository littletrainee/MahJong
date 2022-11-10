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
	tt.FourPair()
	tt.SameKind()
	tt.TwoKang()
	tt.AllPaired()
	tt.TwoDragonHug()
	tt.OneDragon()
	fmt.Println("===============")
	fmt.Printf("共：        %d台\n", tt.total.Get())
}

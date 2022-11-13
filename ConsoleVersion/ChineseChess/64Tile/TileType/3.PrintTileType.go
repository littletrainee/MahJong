package TileType

import "fmt"

func (tt *TileType) PrintTileType() {
	fmt.Println("牌型        台數")
	fmt.Println("---------------")
	tt.Ten_DiOrNinHou()
	tt.TenPai()
	tt.Bookmaker()
	tt.EightGenerals()
	if !tt.iseightgenerals.Get() {
		tt.FiveSoldier()
	}
	tt.ContinueToBookmaker()
	tt.OnlyOrNoGeneralAndSorider()
	tt.TsumoOrRon()
	tt.WinOnLastTile()
	tt.SameKind()
	if tt.eye.Get() != "4pair" {
		if !tt.iseightgenerals.Get() {
			tt.AllPaired()
		}
		tt.TwoKang()
		tt.TwoDragonHug()
		tt.OneDragon()
	} else {
		tt.FourPair()
	}
	fmt.Println("===============")
	fmt.Printf("共：        %d台\n", tt.total.Get())
}

package TileType

import "fmt"

func (tt *TileType) ContinueToBookmaker() {
	if tt.bookmaker.Get() && tt.continuetobookmaker.Get() > 0 {
		v := tt.total.Get()
		v += tt.continuetobookmaker.Get() * 2
		tt.total.Set(v)
		fmt.Printf("連%d拉%d      %d台", tt.continuetobookmaker.Get(),
			tt.continuetobookmaker.Get(), tt.continuetobookmaker.Get()*2)
	}
}

package TileType

import "fmt"

// "天胡"
func (tt *TileType) Ten_DiOrNinHou() {
	if tt.gameround.Get() == 1 {
		if tt.gameturn.Get() == 0 && tt.bookmaker.Get() && tt.istsumo.Get() {
			v := tt.total.Get()
			v += 10
			tt.total.Set(v)
			fmt.Println("天胡       10台")
		} else if tt.gameturn.Get() != 0 && !tt.bookmaker.Get() && tt.istsumo.Get() {
			v := tt.total.Get()
			v += 10
			tt.total.Set(v)
			fmt.Println("地胡       10台")
		} else if tt.gameround.Get() != 0 && !tt.bookmaker.Get() && !tt.istsumo.Get() {
			v := tt.total.Get()
			v += 5
			tt.total.Set(v)
			fmt.Println("人胡        5台")
		}

	}
}

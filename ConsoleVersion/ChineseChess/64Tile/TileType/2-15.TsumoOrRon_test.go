package TileType

import (
	"testing"
)

func TestConcealedAndTsumo(t *testing.T) {
	var (
		tt TileType
	)
	tt.isconcealed.Set(false)
	tt.istsumo.Set(true)
	tt.TsumoOrRon()
}

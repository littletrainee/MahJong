package TileType

import (
	"testing"
)

func TestPrintTileType(t *testing.T) {
	var (
		T TileType
	)
	T.colorlength.Set(1)
	T.kindlength.Set(1)
	T.SetYakuAndPoint()
	T.Print()
}

package TileType

import (
	"github.com/littletrainee/GetSet"
)

type TileType struct {
	yaku        GetSet.Type[string]
	score       GetSet.Type[[]uint8]
	colorlength GetSet.Type[uint8]
	kindlength  GetSet.Type[uint8]
}

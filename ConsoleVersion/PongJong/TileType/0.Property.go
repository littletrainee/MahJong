package TileType

import (
	"github.com/littletrainee/GetSet"
)

type TileType struct {
	target GetSet.Type[string]
	score  GetSet.Type[uint8]
}

package TileType

import (
	"github.com/littletrainee/GetSet"
)

type TileType struct {
	hand      GetSet.Type[[]string]
	total     GetSet.Type[uint8]
	gameround GetSet.Type[uint8]
	gameturn  GetSet.Type[uint8]
	istsumo   GetSet.Type[bool]
	lasttile  GetSet.Type[bool]
	istenpai  GetSet.Type[bool]
}

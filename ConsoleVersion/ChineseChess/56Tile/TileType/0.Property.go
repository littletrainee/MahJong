package TileType

import (
	"github.com/littletrainee/GetSet"
)

type TileType struct {
	hand                GetSet.Type[[]string]
	meld                GetSet.Type[[][]string]
	total               GetSet.Type[uint8]
	gameround           GetSet.Type[uint8]
	gameturn            GetSet.Type[uint8]
	istsumo             GetSet.Type[bool]
	lasttile            GetSet.Type[bool]
	istenpai            GetSet.Type[bool]
	eye                 GetSet.Type[string]
	isconcealed         GetSet.Type[bool]
	twokang             GetSet.Type[uint8]
	winonthewalltail    GetSet.Type[bool]
	bookmaker           GetSet.Type[bool]
	continuetobookmaker GetSet.Type[uint8]
}

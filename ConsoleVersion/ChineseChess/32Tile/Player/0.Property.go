package Player

import (
	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/MahJong/Interface"
)

type Player struct {
	Interface.IPlayer
	Name    GetSet.Type[string]
	Hand    GetSet.Type[[]string]
	River   GetSet.Type[[]string]
	Iswin   GetSet.Type[bool]
	Meld    GetSet.Type[[][]string]
	HasMeld GetSet.Type[[]string]
	IsTsumo GetSet.Type[bool]
	TenPai  GetSet.Type[bool]
}

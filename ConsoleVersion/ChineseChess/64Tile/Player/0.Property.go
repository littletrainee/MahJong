package Player

import (
	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/MahJong/Interface"
)

type Player struct {
	Interface.IPlayer
	Name                GetSet.Type[string]
	Hand                GetSet.Type[[]string]
	River               GetSet.Type[[]string]
	Iswin               GetSet.Type[bool]
	Meld                GetSet.Type[[][]string]
	HasChiMeld          GetSet.Type[[]string]
	HasPongMeld         GetSet.Type[bool]
	HasKangMeld         GetSet.Type[bool]
	IsTsumo             GetSet.Type[bool]
	TenPai              GetSet.Type[bool]
	IsConcealed         GetSet.Type[bool]
	Eye                 GetSet.Type[string]
	TwoKang             GetSet.Type[uint8]
	WinOnTheWallTail    GetSet.Type[bool]
	Bookmaker           GetSet.Type[bool]
	ContinueToBookmaker GetSet.Type[uint8]
	CodeNumber          GetSet.Type[uint8]
}

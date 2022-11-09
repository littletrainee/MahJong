package Player

import (
	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/MahJong/Interface"
)

// func Init() {
// 	rand.Seed(time.Now().UnixNano())
// }

type Player struct {
	Name GetSet.Type[string]
	Hand GetSet.Type[[]uint8]
	Interface.IPlayer
	River GetSet.Type[uint8]
	Iswin GetSet.Type[bool]
}

package Player

import . "github.com/littletrainee/GetSet"

type Player struct {
	Name        Type[string]
	Hand        Type[[]string]
	Meld        Type[[][]string]
	River       Type[[]string]
	CodeNumber  Type[uint8]
	Bookmaker   Type[bool]
	Iswin       Type[bool]
	IsTsumo     Type[bool]
	IsRiiChi    Type[bool]
	HasPongMeld Type[bool]
}

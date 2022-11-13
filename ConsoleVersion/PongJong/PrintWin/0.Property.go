package PrintWin

import (
	"github.com/littletrainee/GetSet"
)

type PrintWin struct {
	name    GetSet.Type[string]
	hand    GetSet.Type[[]string]
	meld    GetSet.Type[[][]string]
	lastone GetSet.Type[string]
	istsumo GetSet.Type[bool]
}

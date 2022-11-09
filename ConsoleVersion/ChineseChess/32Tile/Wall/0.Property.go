package Wall

import (
	"math/rand"
	"time"

	"github.com/littletrainee/GetSet"
	"github.com/littletrainee/MahJong/Interface"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

type Wall struct {
	Name     GetSet.Type[string]
	Wall     GetSet.Type[[]string]
	LastTile GetSet.Type[bool]
	Interface.IWall
}

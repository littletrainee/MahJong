package Player

import (
	"sync"

	"github.com/littletrainee/slices"
)

func (p *Player) CheckBigKang(otherplayerriver string, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := p.Hand.Get()
	if slices.ContainsElement(temp, otherplayerriver) && slices.CountNumber(temp, otherplayerriver) == 3 {
		p.HasKangMeld.Set(true)
	}
}

package Player

import (
	"sync"

	"github.com/littletrainee/slices"
)

func (p *Player) CheckPong(otherplayerriverlastone string, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := p.Hand.Get()
	if slices.ContainsElement(temp, otherplayerriverlastone) && slices.CountNumber(temp, otherplayerriverlastone) == 2 {
		p.HasPongMeld.Set(true)
	}
}

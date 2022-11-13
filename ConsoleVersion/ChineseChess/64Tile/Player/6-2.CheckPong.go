package Player

import (
	"sync"

	"github.com/littletrainee/slices"
)

func (p *Player) CheckPong(otherplayerriver string, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := p.Hand.Get()
	if slices.ContainsElement(temp, otherplayerriver) && slices.CountNumber(temp, otherplayerriver) == 2 {
		p.HasPongMeld.Set(true)
	}
}

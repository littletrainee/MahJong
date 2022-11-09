package Player

import "math/rand"

func (p *Player) DrawCard() {
	temp := p.Hand.Get()
	p.Hand.Set(append(temp, uint8(rand.Intn(7-1)+1)))
}

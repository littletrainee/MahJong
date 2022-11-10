package Player

import "github.com/littletrainee/slices"

func (p *Player) CheckConcealedKang() string {
	for _, v := range p.Hand.Get() {
		if slices.CountNumber(p.Hand.Get(), v) == 4 {
			return v
		}
	}
	return ""
}

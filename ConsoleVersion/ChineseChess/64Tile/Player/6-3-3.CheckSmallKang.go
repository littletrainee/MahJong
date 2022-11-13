package Player

import "github.com/littletrainee/slices"

func (p *Player) CheckSmallKang() string {
	for _, v := range p.Hand.Get() {
		for _, meld := range p.Meld.Get() {
			if slices.CountNumber(meld, v) == 3 {
				return v
			}
		}
	}
	return ""
}

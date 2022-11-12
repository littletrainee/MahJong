package Player

func (p *Player) ResetMeld() {
	p.HasChiMeld.Set([]string{})
	p.HasPongMeld.Set(false)
	p.HasKangMeld.Set(false)
}

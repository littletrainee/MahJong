package Player

// check Tsumo
func (p *Player) TsumoCheck() {
	var (
		iswin    bool
		eye      string
		temphand []string
	)
	// append p's hand to temphand
	temphand = append(temphand, p.Hand.Get()...)

	iswin, eye = endgame(temphand, p.Meld.Get())
	p.Iswin.Set(iswin)
	p.Eye.Set(eye)
}

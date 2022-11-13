package Player

func (p *Player) AskDeclareRiiChi() {
	if asktenpai() == "y" {
		p.IsRiiChi.Set(true)
	}
}

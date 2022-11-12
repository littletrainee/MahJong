package Player

func (p *Player) AskDeclareTenPai() {
	if asktenpai() == "y" {
		p.TenPai.Set(true)
	}
}

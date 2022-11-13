package Player

// Ask Make Meld Or Not
func (p *Player) AskMakeMeldAndChoose() string {
	var (
		key          string
		probablymeld []string = p.HasChiMeld.Get()
	)
	// MakeMeldOrNot:
	key = askMeld(p)
	switch key {
	case "c":
		if len(probablymeld) > 1 { // more than 1 meld
			printprobablychimeld(probablymeld)
			key = askwhichonetochi()
		} else { // only 1 meld
			key = "c1"
		}
	case "p":
		key = "p"
	case "k":
		key = "k"
	case "s":
		key = "s"
	}
	return key
}

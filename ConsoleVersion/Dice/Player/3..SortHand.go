package Player

// Sort Player Hand
func (p *Player) SortHand() {
	// declare variable and Get variable from Player Hand
	var (
		Newtemp []uint8
		temp    []uint8 = p.Hand.Get()
	)
	// check number is in hand, then append to Newtemp
	for i := 1; i < 7; i++ {
		for _, v := range temp {
			if v == uint8(i) {
				Newtemp = append(Newtemp, v)
			}
		}
	}
	// Set Newtemp to player Hand
	p.Hand.Set(Newtemp)
}

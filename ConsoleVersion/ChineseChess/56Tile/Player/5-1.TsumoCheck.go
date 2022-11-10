package Player

// check Tsumo
func (p *Player) TsumoCheck() {
	var (
		iswin      bool
		temphand   []string
		tempplayer Player
	)
	// append p's hand to temphand
	temphand = append(temphand, p.Hand.Get()...)
	// check p's meld isn't empty
	if len(p.Meld.Get()) != 0 {
		temphand = append(temphand, p.Meld.Get()[0]...)
	}
	// sort temphand
	tempplayer.Hand.Set(temphand)
	tempplayer.SortHand()
	temphand = tempplayer.Hand.Get()

	probablyEye := findProbablyEye(temphand)
	for _, v := range probablyEye {
		anothertemp := temphand
		remove_eye_hand := RemoveEye(v, anothertemp)
		meld1 := remove_eye_hand[:3]
		meld2 := remove_eye_hand[3:]
		Ismeld1 := Ismeld(meld1)
		Ismeld2 := Ismeld(meld2)
		if (Ismeld1 == "triple" || Ismeld1 == "sequence") && (Ismeld2 == "triple" || Ismeld2 == "sequence") {
			iswin = true
			p.Eye.Set(v)
			break
		} else {
			iswin = false
		}
	}
	p.Iswin.Set(iswin)
}

package Player

// check Tsumo
func (p *Player) TsumoCheck() {
	var iswin bool
	temp := make([]uint8, len(p.Hand.Get()))
	copy(temp, p.Hand.Get())
	p.SortHand()
	probablyEye := findProbablyEye(temp)
	for _, v := range probablyEye {
		anothertemp := temp
		remove_eye_hand := removeEye(v, anothertemp)
		Ismeld := ismeld(remove_eye_hand)
		if Ismeld == "triple" || Ismeld == "sequence" {
			iswin = true
		} else {
			iswin = false
		}
	}
	p.Iswin.Set(iswin)
}

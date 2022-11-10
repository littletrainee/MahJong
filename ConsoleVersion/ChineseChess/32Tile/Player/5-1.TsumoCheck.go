package Player

import (
	"github.com/littletrainee/slices"
)

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
	if slices.ContainsElement(probablyEye, "1b") && slices.ContainsElement(probablyEye, "1r") {
		anothertemp := temphand
		remove_eye_hand := removeSpecialEye("1b", "1r", anothertemp)
		Ismeld := Ismeld(remove_eye_hand)
		if Ismeld == "triple" || Ismeld == "sequence" {
			// fmt.Printf("\n%s is Tsumo.\n The Hand is %v", p.Name.Get(), p.Hand.Get())
			iswin = true
		} else {
			iswin = false
		}
	} else {
		for _, v := range probablyEye {
			anothertemp := temphand
			remove_eye_hand := removeEye(v, anothertemp)
			Ismeld := Ismeld(remove_eye_hand)
			if Ismeld == "triple" || Ismeld == "sequence" {
				iswin = true
				break
			} else {
				iswin = false
			}
		}
	}
	p.Iswin.Set(iswin)
}

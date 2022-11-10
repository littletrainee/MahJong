package Player

import (
	"sync"

	"github.com/littletrainee/slices"
)

// Check Ron
func (p *Player) RonCheck(otherplayer *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	// declare
	var (
		iswin      bool
		tempplayer Player
		temphand   []string
	)
	// append p's hand to temphand
	temphand = append(temphand, p.Hand.Get()...)
	// append river last to temp
	temphand = append(temphand, otherplayer.River.Get()[len(otherplayer.River.Get())-1])
	// check p's meld isn't empty
	if len(p.Meld.Get()) != 0 {
		temphand = append(temphand, p.Meld.Get()[0]...)
	}
	// Sort temp by tempplayer struct
	tempplayer.Hand.Set(temphand)
	tempplayer.SortHand()
	temphand = tempplayer.Hand.Get()

	// find probably eye
	probablyEye := findProbablyEye(temphand)
	// contain "1b" and "1r"
	if slices.ContainsElement(probablyEye, "1b") && slices.ContainsElement(probablyEye, "1r") {
		anothertemp := temphand
		remove_eye_hand := removeSpecialEye("1b", "1r", anothertemp)
		Ismeld := Ismeld(remove_eye_hand)
		if Ismeld == "triple" || Ismeld == "sequence" {
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
				iswin = false // set tempreturnvalue to false
			}
		}
	}
	p.Iswin.Set(iswin)
}

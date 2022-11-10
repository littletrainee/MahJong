package Player

import (
	"sync"
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
			iswin = false // set tempreturnvalue to false
		}
	}
	p.Iswin.Set(iswin)
}

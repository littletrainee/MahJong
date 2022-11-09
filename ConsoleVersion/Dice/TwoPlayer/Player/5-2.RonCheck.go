package Player

import (
	"sync"
)

// Check Ron
func (p *Player) RonCheck(otherplayer *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		temphand        []uint8
		tempreturnvalue bool
		tempplayer      Player
	)
	// append p's hand to temphand
	temphand = append(temphand, p.Hand.Get()...)
	// append otherplayer's last river to temphand
	temphand = append(temphand, otherplayer.River.Get())

	// Sort temphand by tempplayer struct
	tempplayer.Hand.Set(temphand)
	tempplayer.SortHand()
	temphand = tempplayer.Hand.Get()

	probablyEye := findProbablyEye(temphand)
	for _, v := range probablyEye {
		anothertemp := temphand
		remove_eye_hand := removeEye(v, anothertemp)
		Ismeld := ismeld(remove_eye_hand)
		if Ismeld == "triple" || Ismeld == "sequence" {
			tempreturnvalue = true
		} else {
			tempreturnvalue = false // send NoRon to rt
		}
	}
	p.Iswin.Set(tempreturnvalue)
}

package Player

import (
	"sync"
)

// Check Ron
func (p *Player) RonCheck(otherplayer *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	// declare
	var (
		iswin    bool
		eye      string
		temphand []string
	)
	// append p's hand to temphand
	temphand = append(temphand, p.Hand.Get()...)
	// append river last to temp
	temphand = append(temphand, otherplayer.River.Get()[len(otherplayer.River.Get())-1])

	iswin, eye = endgame(temphand, p.Meld.Get())
	p.Iswin.Set(iswin)
	p.Eye.Set(eye)
}

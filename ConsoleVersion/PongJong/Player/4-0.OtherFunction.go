package Player

// Check meld
func Ismeld(meld []string) bool {
	meld1, meld2, meld3 := meld[0], meld[1], meld[2]
	if meld1 == meld2 && meld2 == meld3 {
		return true
	} else {
		return false
	}
}

func endgame(temphand []string) bool {
	var tempplayer Player
	// sort temphand
	tempplayer.Hand.Set(temphand)
	tempplayer.SortHand()
	temphand = tempplayer.Hand.Get()

	for i := 0; i < 9; i += 3 {
		if !Ismeld([]string{temphand[i], temphand[i+1], temphand[i+2]}) {
			break
		}
		if i == 6 {
			return true
		}
	}
	return false
}

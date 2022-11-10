package Player

import "github.com/littletrainee/slices"

// Find Probably Eye in hand
func findProbablyEye(hand []uint8) []uint8 {
	var probablyeye []uint8
	for i := range hand {
		if i != len(hand)-1 && hand[i] == hand[i+1] {
			if slices.ContainsElement(probablyeye, hand[i]) == false {
				probablyeye = append(probablyeye, hand[i])
			}
		}
	}
	return probablyeye
}

// Remove Eye From Hand
func removeEye(targeteye uint8, orginal []uint8) []uint8 {
	temp := make([]uint8, len(orginal))
	copy(temp, orginal)
	for i := 0; i < 2; i++ {
		index := slices.FindIndexOfElement(temp, targeteye)
		temp = append(temp[:index], temp[index+1:]...)
	}
	return temp
}

// Check meld
func ismeld(meld []uint8) string {
	meld1, meld2, meld3 := meld[0], meld[1], meld[2]
	if meld1 == meld2 && meld2 == meld3 {
		return "triple"
	}
	if meld1+1 == meld2 && meld2+1 == meld3 {
		return "sequence"
	} else {
		return "None"
	}
}

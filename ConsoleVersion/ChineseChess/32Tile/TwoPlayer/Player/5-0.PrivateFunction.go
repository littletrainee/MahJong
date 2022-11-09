package Player

import "github.com/littletrainee/slices"

// Find Probably Eye in hand
func findProbablyEye(hand []string) []string {
	var probablyeye []string
	for i := range hand {
		if i != len(hand)-1 && hand[i] == hand[i+1] {
			if slices.ContainsElement(hand[i], probablyeye) == false {
				probablyeye = append(probablyeye, hand[i])
			}
		}
	}
	// special eye for 32 tile
	if len(probablyeye) == 0 {
		if slices.ContainsElement("1b", hand) && slices.ContainsElement("1r", hand) {
			probablyeye = append(probablyeye, "1b", "1r")
		}
	}
	return probablyeye
}

// Remove Eye From Hand
func removeEye(targeteye string, orginal []string) []string {
	temp := make([]string, len(orginal))
	copy(temp, orginal)
	for i := 0; i < 2; i++ {
		index := slices.FindIndexOfElement(targeteye, temp)
		temp = append(temp[:index], temp[index+1:]...)
	}
	return temp
}

// Remove Special Eye From Hand (only 32 Tile)
func removeSpecialEye(first, second string, orginal []string) []string {
	temp := make([]string, len(orginal))
	copy(temp, orginal)
	index := slices.FindIndexOfElement(first, temp)
	temp = append(temp[:index], temp[index+1:]...)
	index = slices.FindIndexOfElement(second, temp)
	temp = append(temp[:index], temp[index+1:]...)
	return temp
}

// Check meld
func ismeld(meld []string) string {
	meld1, meld2, meld3 := meld[0], meld[1], meld[2]
	if meld1 == meld2 && meld2 == meld3 {
		return "triple"
	}
	if int(meld1[0])+1 == int(meld2[0]) && int(meld2[0])+1 == int(meld3[0]) &&
		meld1[1] == meld2[1] && meld2[1] == meld3[1] {
		return "sequence"
	} else {
		return "None"
	}
}

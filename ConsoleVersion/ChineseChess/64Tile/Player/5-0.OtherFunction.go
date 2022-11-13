package Player

import "github.com/littletrainee/slices"

// Find Probably Eye in hand
func findProbablyEye(hand []string) []string {
	var probablyeye []string
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
func RemoveEye(targeteye string, orginal []string) []string {
	temp := make([]string, len(orginal))
	copy(temp, orginal)
	for i := 0; i < 2; i++ {
		index := slices.FindIndexOfElement(temp, targeteye)
		temp = append(temp[:index], temp[index+1:]...)
	}
	return temp
}

// Remove Special Eye From Hand (only 32 Tile)
func removeSpecialEye(first, second string, orginal []string) []string {
	temp := make([]string, len(orginal))
	copy(temp, orginal)
	index := slices.FindIndexOfElement(temp, first)
	temp = append(temp[:index], temp[index+1:]...)
	index = slices.FindIndexOfElement(temp, second)
	temp = append(temp[:index], temp[index+1:]...)
	return temp
}

// Check meld
func Ismeld(meld []string) string {
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

func isfourpair(temphand []string) bool {
	var returnvalue bool
	// if temphand[0] ==temphand[1]
	for i := 0; i < len(temphand); i += 2 {
		if temphand[i] != temphand[i+1] {
			returnvalue = false
			break
		} else if i == 6 {
			returnvalue = true
		}
	}
	return returnvalue
}

func nomeld(temphand []string) (bool, string) {
	var (
		iswin bool
		eye   string
	)
	if isfourpair(temphand) {
		iswin = true
		eye = "4pair"
	} else {
		// find probably eye
		probablyEye := findProbablyEye(temphand)
		for _, v := range probablyEye {
			anothertemp := make([]string, len(temphand))
			copy(anothertemp, temphand)
			remove_eye_hand := RemoveEye(v, anothertemp)
			meld1 := remove_eye_hand[:3]
			meld2 := remove_eye_hand[3:]
			Ismeld1 := Ismeld(meld1)
			Ismeld2 := Ismeld(meld2)
			if (Ismeld1 == "triple" || Ismeld1 == "sequence") && (Ismeld2 == "triple" || Ismeld2 == "sequence") {
				iswin = true
				eye = v
				break
			} else {
				iswin = false // set tempreturnvalue to false
			}
		}
	}
	return iswin, eye
}

func onemeld(temphand []string) (bool, string) {
	var (
		iswin bool
		eye   string
	)
	// find probably eye
	probablyEye := findProbablyEye(temphand)
	for _, v := range probablyEye {
		anothertemp := make([]string, len(temphand))
		copy(anothertemp, temphand)
		remove_eye_hand := RemoveEye(v, anothertemp)
		ismeld := Ismeld(remove_eye_hand)
		if ismeld == "triple" || ismeld == "sequence" {
			iswin = true
			eye = v
			break
		} else {
			iswin = false // set tempreturnvalue to false
		}
	}
	return iswin, eye
}

func twomeld(temphand []string) (bool, string) {
	return temphand[0] == temphand[1], temphand[0]
}

func endgame(temphand []string, meld [][]string) (bool, string) {
	var tempplayer Player
	// sort temphand
	tempplayer.Hand.Set(temphand)
	tempplayer.SortHand()
	temphand = tempplayer.Hand.Get()

	switch len(meld) {
	case 0: // is concealed
		return nomeld(temphand)
	case 1:
		return onemeld(temphand)
	case 2:
		return twomeld(temphand)
	default:
		return false, ""
	}
}

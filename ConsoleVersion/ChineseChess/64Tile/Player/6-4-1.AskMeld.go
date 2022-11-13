package Player

import (
	"fmt"
	"strings"
)

// Ask Player Want to make meld
func askMeld(player *Player) string {
	var (
		key            string
		probalbyselect string
		c              int  = len(player.HasChiMeld.Get())
		p              bool = player.HasPongMeld.Get()
		k              bool = player.HasKangMeld.Get()
	)

	if c != 0 && p && k { // chi, pong, kang
		fmt.Print("Want to Chi, Pong Or Kang?(c/p/k/s)")
		probalbyselect = "cpks"
	}
	if c != 0 && p && !k { //chi, pong, no kang
		fmt.Print("Want to Chi Or Pong?(c/p/s)")
		probalbyselect = "cps"
	}
	if c != 0 && !p && !k { // chi, no pong, no kang
		fmt.Print("Want to Chi?(c/s)")
		probalbyselect = "cs"
	}
	if c == 0 && p && k { // no chi, pong, kang
		fmt.Print("Want to Pong Or Kang?(p/k/s)")
		probalbyselect = "pks"
	}
	if c == 0 && p && !k { // no chi, pong, no kang
		fmt.Print("Want to Pong?(p/s)")
		probalbyselect = "ps"
	}
	if c == 0 && !p && !k { // no chi, no pong, no kang
		return "s"
	}
	for {
		fmt.Scanln(&key)
		if !strings.Contains(probalbyselect, key) {
			fmt.Printf("%T", key)
			fmt.Print("Wrong Enter Please Renter:")
			key = ""
		} else {
			break
		}
	}
	return key
}

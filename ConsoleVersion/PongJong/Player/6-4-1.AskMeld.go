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
		p              bool = player.HasPongMeld.Get()
	)

	if p { // no chi, pong, no kang
		fmt.Print("Want to Pong?(p/s)")
		probalbyselect = "ps"
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

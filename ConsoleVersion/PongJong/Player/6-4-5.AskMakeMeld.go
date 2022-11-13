package Player

import (
	"fmt"
	"strings"
)

// Ask Make Meld Or Not
func (p *Player) AskMakePong() string {
	var (
		key            string
		probalbyselect string
		pong           bool = p.HasPongMeld.Get()
	)

	if pong { // no chi, pong, no kang
		fmt.Print("Want to Pong?(p/s)")
		probalbyselect = "ps"
	} else {
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

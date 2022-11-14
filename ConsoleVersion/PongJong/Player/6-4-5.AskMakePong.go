package Player

import (
	"fmt"
	"strings"
)

// Ask Make Meld Or Not
func (p *Player) AskMakePong() string {
	var (
		key            string
		probablyoption string
		haspongmeld    bool = p.HasPongMeld.Get()
	)

	if haspongmeld {
		fmt.Print("Want to Pong?(p/s)")
		probablyoption = "ps"
	} else {
		return "s"
	}
	for {
		fmt.Scanln(&key)
		if !strings.Contains(probablyoption, key) {
			fmt.Printf("%T", key)
			fmt.Print("Wrong Enter Please Renter:")
			key = ""
		} else {
			break
		}
	}
	return key
}

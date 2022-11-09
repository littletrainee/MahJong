package Player

import (
	"fmt"
	"strconv"
)

// Discard From Hand To River
func (p *Player) Discard() {
	var (
		keystring string
		keyint    int
	)
	fmt.Print("Please select whitch one do you want to discard from index 1-5:")
	for {
		fmt.Scanf("%v\n", &keystring)
		keyint, _ = strconv.Atoi(keystring)
		if keyint > 5 || keyint <= 0 {
			fmt.Print("Wrong Enter Please Renter:")
			keystring = ""
		} else {
			break
		}
	}
	p.River.Set(p.Hand.Get()[keyint-1])
	p.Hand.Set(append(p.Hand.Get()[:keyint-1], p.Hand.Get()[keyint:]...))
	// p.Hand.Set(temp)
}

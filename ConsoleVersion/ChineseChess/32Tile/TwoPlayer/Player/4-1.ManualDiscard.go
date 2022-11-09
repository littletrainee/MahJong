package Player

import (
	"fmt"
	"strconv"
)

// Discard From Hand To River
func (p *Player) Discard() {
	var (
		keystring string
		temphand  []string = p.Hand.Get()
		tempriver []string = p.River.Get()
	)
	fmt.Printf("Please select whitch one do you want to discard from index 1-%d:", len(p.Hand.Get()))
RenterDiscardChoice:
	fmt.Scanf("%v\n", &keystring)
	keyint, _ := strconv.Atoi(keystring)
	fmt.Printf("the select is: %v\n", keystring)
	for keyint > 5 || keyint <= 0 {
		fmt.Print("Wrong Enter Please Renter:")
		keystring = ""
		goto RenterDiscardChoice
	}
	tempriver = append(tempriver, temphand[keyint-1])
	temphand = append(temphand[:keyint-1], temphand[keyint:]...)
	p.Hand.Set(temphand)
	p.River.Set(tempriver)
}

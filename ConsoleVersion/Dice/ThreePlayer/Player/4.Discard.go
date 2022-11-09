package Player

import (
	"fmt"
	"strconv"
)

// Discard From Hand To River
func (p *Player) Discard() {
	var keystring string
	fmt.Print("Please select whitch one do you want to discard from index 1-5:")
Renter:
	fmt.Scanf("%v\n", &keystring)
	keyint, _ := strconv.Atoi(keystring)
	fmt.Printf("the select is: %v\n", keystring)
	for keyint > 5 || keyint <= 0 {
		fmt.Print("Wrong Enter Please Renter:")
		keystring = ""
		goto Renter
	}
	p.River.Set(p.Hand.Get()[keyint-1])
	temp := append(p.Hand.Get()[:keyint-1], p.Hand.Get()[keyint:]...)
	p.Hand.Set(temp)
}

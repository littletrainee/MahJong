package Player

import "fmt"

func asktenpai() string {
	var key string
	fmt.Printf("Do You Want Declare TenPai?(y/n)")
	fmt.Scanf("%v\n", &key)
RenterDeclareTaiPaiChoice:
	for key != "y" && key != "n" {
		fmt.Print("Wrong Enter Please Renter:")
		key = ""
		goto RenterDeclareTaiPaiChoice
	}
	return key
}
func (p *Player) AskDeclareTenPai() {
	if asktenpai() == "y" {
		p.TenPai.Set(true)
	}
}

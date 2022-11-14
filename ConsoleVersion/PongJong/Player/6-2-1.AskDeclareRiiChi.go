package Player

import "fmt"

func (p *Player) AskDeclareRiiChi() {
	var key string
	fmt.Printf("Do You Want Declare Riichi?(y/n)")
	for {
		fmt.Scanln(&key)
		if key != "y" && key != "n" {
			fmt.Print("Wrong Enter Please Renter:")
			key = ""
		} else {
			break
		}
	}
	if key == "y" {
		p.IsRiiChi.Set(true)
	}
}

package Player

import (
	"fmt"
	"strings"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
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

// print probably meld
func printprobablychimeld(probablymeld []string) {
	fmt.Println()
	for i := 0; i < len(probablymeld); i++ {
		// print index
		fmt.Printf("%d.", i+1)
		CV.PrintRedTextOrNot(probablymeld[i][:2], CC.Tilemap)
		CV.PrintRedTextOrNot(probablymeld[i][2:], CC.Tilemap)
		if i != len(probablymeld)-1 {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
}

// if option more than one ask player which want make meld
func askwhichonetochi() string {
	var key string
	fmt.Printf("Please Select Which One Want Make Meld:")
RenterChoice:
	fmt.Scanln(&key)
	for key != "1" && key != "2" && key != "3" {
		fmt.Print("Wrong Enter Please Renter:")
		key = ""
		goto RenterChoice
	}
	return "c" + key
}

// Ask Make Meld Or Not
func (p *Player) AskMakeMeldAndChoose() string {
	var (
		key          string
		probablymeld []string = p.HasChiMeld.Get()
	)
	// MakeMeldOrNot:
	key = askMeld(p)
	switch key {
	case "c":
		if len(probablymeld) > 1 { // more than 1 meld
			printprobablychimeld(probablymeld)
			key = askwhichonetochi()
		} else { // only 1 meld
			key = "c1"
		}
	case "p":
		key = "p"
	case "k":
		key = "k"
	case "s":
		key = "s"
	}
	return key
}

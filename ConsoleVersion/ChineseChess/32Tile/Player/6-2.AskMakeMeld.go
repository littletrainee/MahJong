package Player

import (
	"fmt"
	"strconv"

	CV "github.com/littletrainee/MahJong/ConsoleVersion"
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
)

// Ask Player Want to make meld
func askMeld() string {
	var key string
	fmt.Printf("Want to make meld?(y/n)")
	for {
		fmt.Scanln(&key)
		if key != "y" && key != "n" {
			fmt.Print("Wrong Enter Please Renter:")
			key = ""
		} else {
			break
		}
	}
	return key
}

// print probably meld
func printprobablymeld(probablymeld []string) {
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
func askwhichone() string {
	var key string
	fmt.Printf("Please Select Which One Want Make Meld:")
RenterChoice:
	fmt.Scanln(&key)
	for key != "1" && key != "2" && key != "3" {
		fmt.Print("Wrong Enter Please Renter:")
		key = ""
		goto RenterChoice
	}
	return key
}

// Ask Make Meld Or Not
func (p *Player) AskMakeMeldAndChoose() int {
	var (
		key          string
		probablymeld []string = p.HasMeld.Get()
		choose       int
	)
	// MakeMeldOrNot:
	key = askMeld()
	if key == "y" {
		if len(probablymeld) > 1 { // more than 1 meld
			printprobablymeld(probablymeld)
			key = askwhichone()
			choose, _ = strconv.Atoi(key)
		} else { // only 1 meld
			choose = 1
		}
	} else if key == "n" {
		choose = 0
	}
	return choose
}

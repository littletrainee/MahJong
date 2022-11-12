package Player

import "fmt"

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

package Player

import "fmt"

func asktenpai() string {
	var key string
	fmt.Printf("Do You Want Declare TenPai?(y/n)")
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

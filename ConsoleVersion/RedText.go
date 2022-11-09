package ConsoleVersion

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/littletrainee/GetSet"
)

var RedString *color.Color = color.New(color.FgRed) // Create new console color with no suffix \n

// Printf red text or orginal color
func PrintRedTextOrNot(s string, Tilemap GetSet.Type[map[string]string]) {
	if s[1] == 'r' {
		RedString.Printf("%v", Tilemap.Get()[s])
	} else {
		fmt.Printf("%v", Tilemap.Get()[s])
	}
}

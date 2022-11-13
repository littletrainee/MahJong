package PrintWin

import "testing"

func TestPrintWinner(t *testing.T) {
	var pw PrintWin
	pw.hand.Set([]string{"bb", "bb", "bb", "bb", "bb", "bb", "bb", "bb"})
	pw.lastone.Set("bb")
	pw.Print()
}

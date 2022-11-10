package Handler

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	var h Handler
	h.Start()
	fmt.Println(h.Player1.Hand.Get())
	fmt.Println(h.Player2.Hand.Get())
}

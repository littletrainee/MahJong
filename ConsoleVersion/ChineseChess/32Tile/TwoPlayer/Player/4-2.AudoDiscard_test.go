package Player

import (
	"fmt"
	"testing"
)

func TestAutoDiscard(t *testing.T) {
	var p Player
	p.Hand.Set([]string{"1b", "2b", "3b", "4b", "5b"})
	p.AutoDiscard()
	if len(p.Hand.Get()) != 4 {
		t.Errorf("Not Success Auto Discard")
	} else {
		fmt.Printf("Success Auto Discard: %v", p.Hand.Get())
	}
}

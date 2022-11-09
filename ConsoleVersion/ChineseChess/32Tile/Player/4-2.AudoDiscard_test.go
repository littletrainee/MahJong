package Player

import (
	"fmt"
	"testing"
)

func TestAutoDiscard(t *testing.T) {
	var p Player
	p.Hand.Set([]string{"1b", "2b"})
	p.AutoDiscard()
	if len(p.Hand.Get()) != 1 {
		t.Errorf("Not Success Auto Discard")
	} else {
		fmt.Printf("Success Auto Discard: %v", p.Hand.Get())
	}
}

package Player

import (
	"sync"
	"testing"
)

func TestRon(t *testing.T) {
	var (
		p1, p2 Player
		wg     *sync.WaitGroup = &sync.WaitGroup{}
	)
	p1.Hand.Set([]string{"7b"})
	p1.Meld.Set([][]string{{"7r", "7r", "7r"}})
	p2.River.Set([]string{"7b"})
	wg.Add(1)
	p1.RonCheck(&p2, wg)
	wg.Wait()
	if p1.Iswin.Get() != true {
		t.Error("Not Win")
	}
}

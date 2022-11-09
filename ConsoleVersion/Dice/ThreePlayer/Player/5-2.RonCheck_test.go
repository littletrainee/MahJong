package Player

import (
	"sync"
	"testing"
)

func TestRonCheck(t *testing.T) {
	var (
		p1, p2 Player
		wg     *sync.WaitGroup = &sync.WaitGroup{}
	)

	p1.Hand.Set([]uint8{1, 2, 3, 4})
	p2.River.Set(1)
	wg.Add(1)
	go p1.RonCheck(&p2, wg)
	wg.Wait()
	if !p1.Iswin.Get() {
		t.Error("Not Win")
	}
}

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
	p1.Hand.Set([]string{"3b", "3b", "2r", "2r", "4r", "4r", "4b"})
	p2.River.Set([]string{"4b"})
	wg.Add(1)
	p1.RonCheck(&p2, wg)
	wg.Wait()
	if p1.Iswin.Get() != true {
		t.Error("Not Win")
	}
}

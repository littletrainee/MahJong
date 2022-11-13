package Player

import (
	"testing"
)

func TestCheckPong(t *testing.T) {
	var (
		p1, p2 Player
		// wg     *sync.WaitGroup = &sync.WaitGroup{}
	)
	p1.Hand.Set([]string{"2b", "2b", "3b", "3b", "3b", "4b", "5b"})
	p2.River.Set([]string{"2b"})
	// if p1.CheckPong(p2.River.Get()[len(p2.River.Get())-1], wg) {
	// 	fmt.Println("yes")
	// } else {
	// 	t.Error("not")
	// }

}

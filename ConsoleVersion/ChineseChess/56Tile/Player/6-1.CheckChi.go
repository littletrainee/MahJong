package Player

import "github.com/littletrainee/slices"

// Check meld in hand
func (p *Player) CheckChi(otherplayerriver string) {
	temp := p.Hand.Get()
	var tempmeld []string
	m2 := string(rune(int(otherplayerriver[0])-2)) + string(otherplayerriver[1])
	m1 := string(rune(int(otherplayerriver[0])-1)) + string(otherplayerriver[1])
	p1 := string(rune(int(otherplayerriver[0])+1)) + string(otherplayerriver[1])
	p2 := string(rune(int(otherplayerriver[0])+2)) + string(otherplayerriver[1])
	if slices.ContainsElement(temp, m2) && slices.ContainsElement(temp, m1) {
		tempmeld = append(tempmeld, m2+m1)
	}
	if slices.ContainsElement(temp, m1) && slices.ContainsElement(temp, p1) {
		tempmeld = append(tempmeld, m1+p1)
	}
	if slices.ContainsElement(temp, p1) && slices.ContainsElement(temp, p2) {
		tempmeld = append(tempmeld, p1+p2)
	}
	p.HasChiMeld.Set(tempmeld)
}

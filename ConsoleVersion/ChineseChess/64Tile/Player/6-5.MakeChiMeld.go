package Player

import (
	"github.com/littletrainee/slices"
)

// Make meld
func (p *Player) MakeChiMeld(choice string, otherplayer *Player) {
	var (
		temphand         []string   = p.Hand.Get()
		tempmeldlist     [][]string = p.Meld.Get()
		tempmeld         []string
		otherplayerriver []string = otherplayer.River.Get()
		tempmeldtarget   string   = p.HasChiMeld.Get()[int(choice[1]-49)]
		index            int
	)
	// append meld to tempmeld
	tempmeld = append(tempmeld, tempmeldtarget[:2])                 // add first element of hand to tempmeld
	index = slices.FindIndexOfElement(temphand, tempmeldtarget[:2]) //find first element index in hand
	temphand = append(temphand[:index], temphand[index+1:]...)      // remove first element from hand

	tempmeld = append(tempmeld, otherplayerriver[len(otherplayerriver)-1]) // add last element of river to tempmeld
	otherplayerriver = otherplayerriver[:len(otherplayerriver)-1]          // remove last element form river

	tempmeld = append(tempmeld, tempmeldtarget[2:])                 // add second element of hand to tempmeld
	index = slices.FindIndexOfElement(temphand, tempmeldtarget[2:]) // find second element index in hand
	temphand = append(temphand[:index], temphand[index+1:]...)      // remove second element from hand

	tempmeldlist = append(tempmeldlist, tempmeld) // Add tempmeld to tempmeldlist

	p.HasChiMeld.Set([]string{})            // clear hasmeld
	otherplayer.River.Set(otherplayerriver) // Set other player.river
	p.Hand.Set(temphand)                    // Set temphand to Hand
	p.Meld.Set(tempmeldlist)                // Set tempmeldlist to Meld
	p.IsConcealed.Set(false)
}

package Player

import (
	CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"
	"github.com/littletrainee/slices"
)

func (p *Player) TenPaiCheck() []string {
	var (
		temphand           []string = p.Hand.Get()
		tempplayer         Player
		cloneforpopone     []string
		removeonefromclone []string
		probablywintile    []string
	)

	// Sort temphand
	tempplayer.Hand.Set(temphand)
	tempplayer.SortHand()
	temphand = tempplayer.Hand.Get()

	// pop one by one to check is tenpai
	for i := range temphand {
		// reset cloneforpopone capacity and value
		cloneforpopone = make([]string, len(temphand))
		// clone from temphand to cloneforpopone[5]
		copy(cloneforpopone, temphand)
		// remove one from cloneforpopone array by temphand index [5->4]
		cloneforpopone = append(cloneforpopone[:i], cloneforpopone[i+1:]...)

		// try add one to array[4->5]
		for _, probablytarget := range CC.Tile_key {
			// reset removeonefromclone capacity and value
			removeonefromclone = make([]string, len(cloneforpopone))
			// clone from cloneforpopone to removeoonefromclone[4]
			copy(removeonefromclone, cloneforpopone)
			removeonefromclone = append(removeonefromclone, probablytarget)
			if probablywin(removeonefromclone, p.Meld.Get()) {
				if !slices.ContainsElement(probablywintile, probablytarget) {
					probablywintile = append(probablywintile, probablytarget)
				}
			}
		}
	}
	return probablywintile
}

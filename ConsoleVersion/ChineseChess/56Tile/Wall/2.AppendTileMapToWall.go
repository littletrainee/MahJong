package Wall

import CC "github.com/littletrainee/MahJong/ConsoleVersion/ChineseChess"

// append tile map to wall
func (w *Wall) AppendTileMapToWall() {
	var temp []string
	for i := range CC.Tilemap.Get() {
		for j := 0; j < 4; j++ {
			temp = append(temp, i)
		}
	}
	w.Wall.Set(temp)
}
